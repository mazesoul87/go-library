// Copyright 2021-present The Atlas Authors. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sqltool

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"
	"unicode"

	"ariga.io/atlas/sql/migrate"
)

var (
	// GolangMigrateFormatter returns migrate.Formatter compatible with golang-migrate/migrate.
	GolangMigrateFormatter = templateFormatter(
		"{{ now }}{{ with .Name }}_{{ . }}{{ end }}.up.sql",
		`{{ range .Changes }}{{ with .Comment }}-- {{ println . }}{{ end }}{{ printf "%s;\n" .Cmd }}{{ end }}`,
		"{{ now }}{{ with .Name }}_{{ . }}{{ end }}.down.sql",
		`{{ range $c := rev .Changes }}{{ with $stmts := .ReverseStmts }}{{ with $c.Comment }}-- reverse: {{ println . }}{{ end }}{{ range $stmts }}{{ printf "%s;\n" . }}{{ end }}{{ end }}{{ end }}`,
	)
	// GooseFormatter returns migrate.Formatter compatible with pressly/goose.
	GooseFormatter = templateFormatter(
		"{{ now }}{{ with .Name }}_{{ . }}{{ end }}.sql",
		`-- +goose Up
{{ range .Changes }}{{ with .Comment }}-- {{ println . }}{{ end }}{{ printf "%s;\n" .Cmd }}{{ end }}
-- +goose Down
{{ range $c := rev .Changes }}{{ with $stmts := .ReverseStmts }}{{ with $c.Comment }}-- reverse: {{ println . }}{{ end }}{{ range $stmts }}{{ printf "%s;\n" . }}{{ end }}{{ end }}{{ end }}`,
	)
	// FlywayFormatter returns migrate.Formatter compatible with Flyway.
	FlywayFormatter = templateFormatter(
		"V{{ now }}{{ with .Name }}__{{ . }}{{ end }}.sql",
		`{{ range .Changes }}{{ with .Comment }}-- {{ println . }}{{ end }}{{ printf "%s;\n" .Cmd }}{{ end }}`,
		"U{{ now }}{{ with .Name }}__{{ . }}{{ end }}.sql",
		`{{ range $c := rev .Changes }}{{ with $stmts := .ReverseStmts }}{{ with $c.Comment }}-- reverse: {{ println . }}{{ end }}{{ range $stmts }}{{ printf "%s;\n" . }}{{ end }}{{ end }}{{ end }}`,
	)
	// LiquibaseFormatter returns migrate.Formatter compatible with Liquibase.
	LiquibaseFormatter = templateFormatter(
		"{{ now }}{{ with .Name }}_{{ . }}{{ end }}.sql",
		`{{- $now := now -}}
--liquibase formatted sql

{{- range $index, $change := .Changes }}
--changeset atlas:{{ $now }}-{{ inc $index }}
{{ with $change.Comment }}--comment: {{ . }}{{ end }}
{{ $change.Cmd }};
{{ with $stmts := .ReverseStmts }}{{ range $stmts }}{{ printf "--rollback: %s;\n" . }}{{ end }}{{ end }}
{{- end }}`,
	)
	// DBMateFormatter returns migrate.Formatter compatible with amacneil/dbmate.
	DBMateFormatter = templateFormatter(
		"{{ now }}{{ with .Name }}_{{ . }}{{ end }}.sql",
		`-- migrate:up
{{ range .Changes }}{{ with .Comment }}-- {{ println . }}{{ end }}{{ printf "%s;\n" .Cmd }}{{ end }}
-- migrate:down
{{ range $c := rev .Changes }}{{ with $stmts := .ReverseStmts }}{{ with $c.Comment }}-- reverse: {{ println . }}{{ end }}{{ range $stmts }}{{ printf "%s;\n" . }}{{ end }}{{ end }}{{ end }}`,
	)
	// DbmateFormatter is the same as DBMateFormatter.
	// Deprecated: Use DBMateFormatter instead.
	DbmateFormatter = DBMateFormatter
)

type (
	// GolangMigrateDir wraps fs.FS and provides a migrate.Scanner implementation able to understand files
	// generated by the GolangMigrateFormatter for migration directory replaying.
	GolangMigrateDir struct{ fs.FS }
	// GolangMigrateFile wraps migrate.LocalFile with custom description function.
	GolangMigrateFile struct{ *migrate.LocalFile }
)

// NewGolangMigrateDir returns a new GolangMigrateDir.
func NewGolangMigrateDir(path string) (*GolangMigrateDir, error) {
	dir, err := migrate.NewLocalDir(path)
	if err != nil {
		return nil, err
	}
	return &GolangMigrateDir{dir}, nil
}

// Path returns the local path used for opening this dir.
func (d *GolangMigrateDir) Path() string {
	if dir, ok := d.FS.(dirPath); ok {
		return dir.Path()
	}
	return ""
}

// Files implements Scanner.Files. It looks for all files with up.sql suffix and orders them by filename.
func (d *GolangMigrateDir) Files() ([]migrate.File, error) {
	names, err := fs.Glob(d, "*.up.sql")
	if err != nil {
		return nil, err
	}
	// Sort files lexicographically.
	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})
	ret := make([]migrate.File, len(names))
	for i, n := range names {
		b, err := fs.ReadFile(d, n)
		if err != nil {
			return nil, fmt.Errorf("sql/migrate: read file %q: %w", n, err)
		}
		ret[i] = &GolangMigrateFile{LocalFile: migrate.NewLocalFile(n, b)}
	}
	return ret, nil
}

// Checksum implements Dir.Checksum. By default, it calls Files() and creates a checksum from them.
func (d *GolangMigrateDir) Checksum() (migrate.HashFile, error) {
	if d, ok := d.FS.(migrate.Dir); ok {
		return d.Checksum()
	}
	files, err := d.Files()
	if err != nil {
		return nil, err
	}
	return migrate.NewHashFile(files)
}

// WriteFile implements Dir.WriteFile.
func (d *GolangMigrateDir) WriteFile(name string, b []byte) error {
	if d, ok := d.FS.(migrate.Dir); ok {
		return d.WriteFile(name, b)
	}
	return errors.New("sql/sqltool: write not supported")
}

// Desc implements File.Desc.
func (f *GolangMigrateFile) Desc() string {
	return strings.TrimSuffix(f.LocalFile.Desc(), ".up")
}

type (
	// GooseDir wraps migrate.LocalDir and provides a migrate.Scanner implementation able to understand files
	// generated by the GooseFormatter for migration directory replaying.
	GooseDir struct{ *migrate.LocalDir }
	// GooseFile wraps migrate.LocalFile with custom statements function.
	GooseFile struct{ *migrate.LocalFile }
)

// NewGooseDir returns a new GooseDir.
func NewGooseDir(path string) (*GooseDir, error) {
	dir, err := migrate.NewLocalDir(path)
	if err != nil {
		return nil, err
	}
	return &GooseDir{dir}, nil
}

// Files looks for all files with .sql suffix and orders them by filename.
func (d *GooseDir) Files() ([]migrate.File, error) {
	files, err := d.LocalDir.Files()
	if err != nil {
		return nil, err
	}
	for i, f := range files {
		files[i] = &GooseFile{f.(*migrate.LocalFile)}
	}
	return files, nil
}

// StmtDecls understands the migration format used by pressly/goose sql migration files.
func (f *GooseFile) StmtDecls() ([]*migrate.Stmt, error) {
	// Atlas custom delimiter is per file, goose has pragma do mark start and end of a delimiter.
	// In order to use the Atlas lexer, we define a custom delimiter for the source SQL and edit it to use the
	// custom delimiter.
	const delim = "-- ATLAS_DELIM_END"
	var (
		state, lineCount int
		lines            = []string{"-- atlas:delimiter " + delim, ""}
		sc               = bufio.NewScanner(bytes.NewReader(f.Bytes()))
	)
Scan:
	for sc.Scan() {
		lineCount++
		line := sc.Text()
		// Handle goose custom delimiters.
		if strings.HasPrefix(line, goosePragma) {
			switch strings.TrimSpace(strings.TrimPrefix(line, goosePragma)) {
			case "Up":
				switch state {
				case none: // found the "up" part of the file
					state = up
				default:
					return nil, unexpectedPragmaErr(f, lineCount, "Up")
				}
			case "Down":
				switch state {
				case up: // found the "down" part
					break Scan
				default:
					return nil, unexpectedPragmaErr(f, lineCount, "Down")
				}
			case "StatementBegin":
				switch state {
				case up:
					state = begin // begin of a statement
				default:
					return nil, unexpectedPragmaErr(f, lineCount, "StatementBegin")
				}
			case "StatementEnd":
				switch state {
				case begin:
					state = end // end of a statement
				default:
					return nil, unexpectedPragmaErr(f, lineCount, "StatementEnd")
				}
			}
		}
		// Write the line of the statement.
		if !reGoosePragma.MatchString(line) && state != end {
			// end of statement if line ends with semicolon
			line = strings.TrimRightFunc(line, unicode.IsSpace)
			lines = append(lines, line)
			if state == up && strings.HasSuffix(line, ";") && !strings.HasPrefix(line, "--") {
				lines = append(lines, delim)
			}
		}
		if state == end {
			state = up
			lines = append(lines, delim)
		}
	}
	return migrate.Stmts(strings.Join(lines, "\n"))
}

// Stmts understands the migration format used by pressly/goose sql migration files.
func (f *GooseFile) Stmts() ([]string, error) {
	s, err := f.StmtDecls()
	if err != nil {
		return nil, err
	}
	stmts := make([]string, len(s))
	for i := range s {
		stmts[i] = s[i].Text
	}
	return stmts, nil
}

type (
	// DBMateDir wraps migrate.LocalDir and provides a migrate.Scanner implementation able to understand files
	// generated by the DBMateFormatter for migration directory replaying.
	DBMateDir struct{ *migrate.LocalDir }
	// DBMateFile wraps migrate.LocalFile with custom statements function.
	DBMateFile struct{ *migrate.LocalFile }
)

// NewDBMateDir returns a new DBMateDir.
func NewDBMateDir(path string) (*DBMateDir, error) {
	dir, err := migrate.NewLocalDir(path)
	if err != nil {
		return nil, err
	}
	return &DBMateDir{dir}, nil
}

// Files looks for all files with up.sql suffix and orders them by filename.
func (d *DBMateDir) Files() ([]migrate.File, error) {
	files, err := d.LocalDir.Files()
	if err != nil {
		return nil, err
	}
	for i, f := range files {
		files[i] = &DBMateFile{f.(*migrate.LocalFile)}
	}
	return files, nil
}

// StmtDecls understands the migration format used by amacneil/dbmate sql migration files.
func (f *DBMateFile) StmtDecls() ([]*migrate.Stmt, error) {
	var (
		state, lineCount int
		lines            []string
		sc               = bufio.NewScanner(bytes.NewReader(f.Bytes()))
	)
Scan:
	for sc.Scan() {
		lineCount++
		line := sc.Text()
		// Handle pragmas.
		if strings.HasPrefix(line, dbmatePragma) {
			switch strings.TrimSpace(strings.TrimPrefix(line, dbmatePragma)) {
			case "up":
				state = up
			case "down":
				break Scan
			}
		}
		// Write the line of the statement.
		if !reDBMatePragma.MatchString(line) && state == up {
			lines = append(lines, line)
		}
	}
	return migrate.Stmts(strings.Join(lines, "\n"))
}

// Stmts understands the migration format used by amacneil/dbmate sql migration files.
func (f *DBMateFile) Stmts() ([]string, error) {
	s, err := f.StmtDecls()
	if err != nil {
		return nil, err
	}
	stmts := make([]string, len(s))
	for i := range s {
		stmts[i] = s[i].Text
	}
	return stmts, nil
}

type (
	// FlywayDir wraps fs.FS and provides a migrate.Scanner implementation able to understand files
	// generated by the FlywayFormatter for migration directory replaying.
	FlywayDir struct{ fs.FS }
	// FlywayFile wraps migrate.LocalFile with custom statements function.
	FlywayFile struct{ *migrate.LocalFile }
)

// NewFlywayDir returns a new FlywayDir.
func NewFlywayDir(path string) (*FlywayDir, error) {
	dir, err := migrate.NewLocalDir(path)
	if err != nil {
		return nil, err
	}
	return &FlywayDir{dir}, nil
}

// Path returns the local path used for opening this dir.
func (d *FlywayDir) Path() string {
	if dir, ok := d.FS.(dirPath); ok {
		return dir.Path()
	}
	return ""
}

// Files implements Scanner.Files. It looks for all files with .sql suffix. The given directory is recursively scanned
// for non-hidden subdirectories. All found files will be ordered by migration type (Baseline, Versioned, Repeatable)
// and filename.
func (d *FlywayDir) Files() ([]migrate.File, error) {
	var ff flywayFiles
	if err := fs.WalkDir(d, ".", func(path string, e fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path != "." && e.IsDir() {
			fullPath := path
			if p, ok := d.FS.(dirPath); ok {
				fullPath = filepath.Join(p.Path(), path)
			}
			h, err := hidden(fullPath)
			if err != nil {
				return err
			}
			if h {
				return fs.SkipDir
			}
			return nil
		}
		var (
			pfx  = e.Name()[0]
			base = filepath.Base(e.Name())
			ext  = filepath.Ext(e.Name())
		)
		if ext != ".sql" || len(base) < 4 || (pfx != 'V' && pfx != 'B' && pfx != 'R') {
			return nil
		}
		return ff.add(path)
	}); err != nil {
		return nil, err
	}
	var (
		names = ff.names()
		ret   = make([]migrate.File, len(names))
	)
	for i, n := range names {
		b, err := fs.ReadFile(d, n)
		if err != nil {
			return nil, fmt.Errorf("sql/migrate: read file %q: %w", n, err)
		}
		ret[i] = &FlywayFile{migrate.NewLocalFile(n, b)}
	}
	return ret, nil
}

// Checksum implements Dir.Checksum. By default, it calls Files() and creates a checksum from them.
func (d *FlywayDir) Checksum() (migrate.HashFile, error) {
	if d, ok := d.FS.(migrate.Dir); ok {
		return d.Checksum()
	}
	files, err := d.Files()
	if err != nil {
		return nil, err
	}
	return migrate.NewHashFile(files)
}

// WriteFile implements Dir.WriteFile.
func (d *FlywayDir) WriteFile(name string, b []byte) error {
	if d, ok := d.FS.(migrate.Dir); ok {
		return d.WriteFile(name, b)
	}
	return errors.New("sql/sqltool: write not supported")
}

// Desc implements File.Desc.
func (f FlywayFile) Desc() string {
	return flywayDesc(f.Name())
}

// Version implements File.Version.
func (f FlywayFile) Version() string {
	return flywayVersion(f.Name())
}

// SetRepeatableVersion iterates over the migration files and assigns repeatable migrations a version number since
// Atlas does not have the concept of repeatable migrations. Each repeatable migration file gets assigned the version
// of the preceding migration file (or 0) followed by an 'R'.
func SetRepeatableVersion(ff []migrate.File) {
	// First find the index of the first repeatable migration file (if any).
	var (
		v   string // last versioned migration version
		idx = func() int {
			for i, f := range ff {
				if f.Version() == "" {
					return i
				}
			}
			return -1
		}()
	)
	switch idx {
	case -1:
		// No repeatable migration does exist.
		return
	case 0:
		// There is no preceding migration. Use Version "0".
		v = "0"
	default:
		v = ff[idx-1].Version()
	}
	if v != "" {
		// Every migration file following the first repeatable found are repeatable as well.
		for i, f := range ff[idx:] {
			ff[idx+i] = &FlywayFile{migrate.NewLocalFile(
				fmt.Sprintf("V%sR__%s", v, f.Desc()),
				f.Bytes(),
			)}
		}
	}
}

// LiquibaseDir wraps migrate.LocalDir and provides a migrate.Scanner implementation able to understand files
// generated by the LiquibaseFormatter for migration directory replaying.
type LiquibaseDir struct{ *migrate.LocalDir }

// NewLiquibaseDir returns a new LiquibaseDir.
func NewLiquibaseDir(path string) (*LiquibaseDir, error) {
	d, err := migrate.NewLocalDir(path)
	if err != nil {
		return nil, err
	}
	return &LiquibaseDir{d}, nil
}

const (
	none int = iota
	up
	begin
	end
	goosePragma  = "-- +goose"
	dbmatePragma = "-- migrate:"
)

var (
	reGoosePragma  = regexp.MustCompile(regexp.QuoteMeta(goosePragma) + " Up|Down|StatementBegin|StatementEnd")
	reDBMatePragma = regexp.MustCompile(dbmatePragma + "up|down")
)

// flywayFiles retrieves flyway migration files by calls to add(). It will only keep the latest baseline and ignore
// all versioned files that are included in that baseline.
type flywayFiles struct {
	baseline   string
	versioned  []string
	repeatable []string
}

// add the given path to the migration files according to its type. The input directory is assumed to be valid
// according to the Flyway documentation (no duplicate versions, etc.).
func (ff *flywayFiles) add(path string) error {
	switch p := filepath.Base(path)[0]; p {
	case 'B':
		if ff.baseline != "" && flywayVersion(path) < flywayVersion(ff.baseline) {
			return nil
		}
		ff.baseline = path
		// In case we set a new baseline, remove all versioned files with a version smaller than the new baseline.
		var (
			bv = flywayVersion(ff.baseline)
			vs []string
		)
		for _, v := range ff.versioned {
			if v > bv {
				vs = append(vs, v)
			}
		}
		ff.versioned = vs
		return nil
	case 'V':
		v := flywayVersion(path)
		if ff.baseline == "" || flywayVersion(ff.baseline) < v {
			ff.versioned = append(ff.versioned, path)
		}
		return nil
	case 'R':
		ff.repeatable = append(ff.repeatable, path)
		return nil
	default:
		return fmt.Errorf("sql/sqltool: unexpected Flyway prefix %q", p)
	}
}

func (ff *flywayFiles) names() []string {
	var names []string
	if ff.baseline != "" {
		names = append(names, ff.baseline)
	}
	flywaySort(ff.versioned)
	flywaySort(ff.repeatable)
	names = append(names, ff.versioned...)
	names = append(names, ff.repeatable...)
	return names
}

func flywayDesc(path string) string {
	parts := strings.SplitN(path, "__", 2)
	if len(parts) == 1 {
		return ""
	}
	return strings.TrimSuffix(parts[1], ".sql")
}

func flywayVersion(path string) string {
	// Repeatable migrations don't have a version.
	if filepath.Base(path)[0] == 'R' {
		return ""
	}
	return strings.SplitN(strings.TrimSuffix(filepath.Base(path), ".sql"), "__", 2)[0][1:]
}

func flywaySort(files []string) {
	sort.Slice(files, func(i, j int) bool {
		return flywayVersionCompare(flywayVersion(files[i]), flywayVersion(files[j])) < 0
	})
}

func flywayVersionCompare(v1, v2 string) int {
	parse := func(s string) []int {
		ss := strings.Split(strings.ReplaceAll(s, "_", "."), ".")
		var ret []int
		for _, s := range ss {
			// 0 for non-numeric parts.
			i, _ := strconv.Atoi(s)
			ret = append(ret, i)
		}
		return ret
	}
	return compareSliceInt(parse(v1), parse(v2))
}

func unexpectedPragmaErr(f migrate.File, line int, pragma string) error {
	var tool string
	switch f := f.(type) {
	case *GooseFile:
		tool = "goose"
	case *DBMateFile:
		tool = "dbmate"
	default:
		return fmt.Errorf("sql/migrate: unexpected migration file type '%T'", f)
	}
	return fmt.Errorf(
		"sql/migrate: %s: %s:%d unexpected goosePragma '%s'",
		tool, f.Name(), line, pragma,
	)
}

// funcs contains the template.FuncMap for the different formatters.
var funcs = template.FuncMap{
	"inc": func(x int) int { return x + 1 },
	// now formats the current time in a lexicographically ascending order while maintaining human readability.
	"now": func() string { return time.Now().UTC().Format("20060102150405") },
	"rev": reverse,
}

// templateFormatter parses the given templates and passes them on to the migrate.NewTemplateFormatter.
func templateFormatter(templates ...string) migrate.Formatter {
	tpls := make([]*template.Template, len(templates))
	for i, t := range templates {
		tpls[i] = template.Must(template.New("").Funcs(funcs).Parse(t))
	}
	tf, err := migrate.NewTemplateFormatter(tpls...)
	if err != nil {
		panic(err)
	}
	return tf
}

// reverse changes for the down migration.
func reverse(changes []*migrate.Change) []*migrate.Change {
	n := len(changes)
	rev := make([]*migrate.Change, n)
	if n%2 == 1 {
		rev[n/2] = changes[n/2]
	}
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = changes[j], changes[i]
	}
	return rev
}

type dirPath interface {
	Path() string
}

var (
	_ dirPath = (*DBMateDir)(nil)
	_ dirPath = (*FlywayDir)(nil)
	_ dirPath = (*GolangMigrateDir)(nil)
	_ dirPath = (*GooseDir)(nil)
	_ dirPath = (*LiquibaseDir)(nil)
)

// Copied from golang.org/x/exp/slices.Compare
func compareSliceInt(s1, s2 []int) int {
	s2len := len(s2)
	for i, v1 := range s1 {
		if i >= s2len {
			return +1
		}
		v2 := s2[i]
		switch {
		case v1 < v2:
			return -1
		case v1 > v2:
			return +1
		}
	}
	if len(s1) < s2len {
		return -1
	}
	return 0
}
