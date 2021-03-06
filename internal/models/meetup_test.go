// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testMeetups(t *testing.T) {
	t.Parallel()

	query := Meetups()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMeetupsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meetup{}
	if err = randomize.Struct(seed, o, meetupDBTypes, true, meetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Meetups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMeetupsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meetup{}
	if err = randomize.Struct(seed, o, meetupDBTypes, true, meetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Meetups().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Meetups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMeetupsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meetup{}
	if err = randomize.Struct(seed, o, meetupDBTypes, true, meetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MeetupSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Meetups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMeetupsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meetup{}
	if err = randomize.Struct(seed, o, meetupDBTypes, true, meetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MeetupExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Meetup exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MeetupExists to return true, but got false.")
	}
}

func testMeetupsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meetup{}
	if err = randomize.Struct(seed, o, meetupDBTypes, true, meetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	meetupFound, err := FindMeetup(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if meetupFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMeetupsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meetup{}
	if err = randomize.Struct(seed, o, meetupDBTypes, true, meetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Meetups().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMeetupsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meetup{}
	if err = randomize.Struct(seed, o, meetupDBTypes, true, meetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Meetups().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMeetupsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	meetupOne := &Meetup{}
	meetupTwo := &Meetup{}
	if err = randomize.Struct(seed, meetupOne, meetupDBTypes, false, meetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}
	if err = randomize.Struct(seed, meetupTwo, meetupDBTypes, false, meetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = meetupOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = meetupTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Meetups().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMeetupsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	meetupOne := &Meetup{}
	meetupTwo := &Meetup{}
	if err = randomize.Struct(seed, meetupOne, meetupDBTypes, false, meetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}
	if err = randomize.Struct(seed, meetupTwo, meetupDBTypes, false, meetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = meetupOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = meetupTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Meetups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func meetupBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Meetup) error {
	*o = Meetup{}
	return nil
}

func meetupAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Meetup) error {
	*o = Meetup{}
	return nil
}

func meetupAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Meetup) error {
	*o = Meetup{}
	return nil
}

func meetupBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Meetup) error {
	*o = Meetup{}
	return nil
}

func meetupAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Meetup) error {
	*o = Meetup{}
	return nil
}

func meetupBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Meetup) error {
	*o = Meetup{}
	return nil
}

func meetupAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Meetup) error {
	*o = Meetup{}
	return nil
}

func meetupBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Meetup) error {
	*o = Meetup{}
	return nil
}

func meetupAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Meetup) error {
	*o = Meetup{}
	return nil
}

func testMeetupsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Meetup{}
	o := &Meetup{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, meetupDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Meetup object: %s", err)
	}

	AddMeetupHook(boil.BeforeInsertHook, meetupBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	meetupBeforeInsertHooks = []MeetupHook{}

	AddMeetupHook(boil.AfterInsertHook, meetupAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	meetupAfterInsertHooks = []MeetupHook{}

	AddMeetupHook(boil.AfterSelectHook, meetupAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	meetupAfterSelectHooks = []MeetupHook{}

	AddMeetupHook(boil.BeforeUpdateHook, meetupBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	meetupBeforeUpdateHooks = []MeetupHook{}

	AddMeetupHook(boil.AfterUpdateHook, meetupAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	meetupAfterUpdateHooks = []MeetupHook{}

	AddMeetupHook(boil.BeforeDeleteHook, meetupBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	meetupBeforeDeleteHooks = []MeetupHook{}

	AddMeetupHook(boil.AfterDeleteHook, meetupAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	meetupAfterDeleteHooks = []MeetupHook{}

	AddMeetupHook(boil.BeforeUpsertHook, meetupBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	meetupBeforeUpsertHooks = []MeetupHook{}

	AddMeetupHook(boil.AfterUpsertHook, meetupAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	meetupAfterUpsertHooks = []MeetupHook{}
}

func testMeetupsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meetup{}
	if err = randomize.Struct(seed, o, meetupDBTypes, true, meetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Meetups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMeetupsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meetup{}
	if err = randomize.Struct(seed, o, meetupDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(meetupColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Meetups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMeetupToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Meetup
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, meetupDBTypes, false, meetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := MeetupSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Meetup)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testMeetupToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Meetup
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, meetupDBTypes, false, strmangle.SetComplement(meetupPrimaryKeyColumns, meetupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Meetups[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testMeetupsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meetup{}
	if err = randomize.Struct(seed, o, meetupDBTypes, true, meetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMeetupsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meetup{}
	if err = randomize.Struct(seed, o, meetupDBTypes, true, meetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MeetupSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMeetupsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Meetup{}
	if err = randomize.Struct(seed, o, meetupDBTypes, true, meetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Meetups().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	meetupDBTypes = map[string]string{`ID`: `bigint`, `Names`: `character varying`, `Description`: `character varying`, `UserID`: `bigint`}
	_             = bytes.MinRead
)

func testMeetupsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(meetupPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(meetupAllColumns) == len(meetupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Meetup{}
	if err = randomize.Struct(seed, o, meetupDBTypes, true, meetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Meetups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, meetupDBTypes, true, meetupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMeetupsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(meetupAllColumns) == len(meetupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Meetup{}
	if err = randomize.Struct(seed, o, meetupDBTypes, true, meetupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Meetups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, meetupDBTypes, true, meetupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(meetupAllColumns, meetupPrimaryKeyColumns) {
		fields = meetupAllColumns
	} else {
		fields = strmangle.SetComplement(
			meetupAllColumns,
			meetupPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := MeetupSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMeetupsUpsert(t *testing.T) {
	t.Parallel()

	if len(meetupAllColumns) == len(meetupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Meetup{}
	if err = randomize.Struct(seed, &o, meetupDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Meetup: %s", err)
	}

	count, err := Meetups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, meetupDBTypes, false, meetupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Meetup struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Meetup: %s", err)
	}

	count, err = Meetups().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
