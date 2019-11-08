package _generated

//go:generate msgp

type DuplicatedFieldTag struct {
	Field1 int    `msg:"t"`
	Field2 string `msg:"t"`
}

type DuplicatedFieldTag2 struct {
	EmbeddableStruct `msg:",flatten"`
	Field1           int `msg:"t"`
}

type EmbeddedStrut struct {
	Field2 string `msg:"t"`
}
