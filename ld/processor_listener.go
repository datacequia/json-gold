package ld

type ProcessorListener interface {
	OnTermDefinition(TermDefinition)
}

type TermDefinition interface {
	Term() string
	Value() TermValue
}

type TermValue struct {
	ReverseKeyword bool
	IdKeyword      string
	IsPrefix       bool // IdKeyword is a prefix IRI
	IsProtected    bool
	TypeKeyword    string
	//ContainerKeyword interface{}
	//ContextKeyword   interface{}
	LanguageKeyword  string
	DirectionKeyword string
	NestKeyword      string
	IndexKeyword     string
}

type termDefinition struct {
	term  string
	value TermValue
}

func newTermDefinition(term string, value TermValue) *termDefinition {

	return &termDefinition{
		term:  term,
		value: value,
	}

}

func (td *termDefinition) Term() string {
	return td.term
}

func (td *termDefinition) Value() TermValue {

	return td.value

}

func (tv *TermValue) copy(termAttrs map[string]interface{}) {

	for termKey, termValue := range termAttrs {

		switch termKey {
		case "@reverse":
			if v, ok := termValue.(bool); ok {
				tv.ReverseKeyword = v
			}

		case "@id":
			//tv.IdKeyword:
			if v, ok := termValue.(string); ok {
				tv.IdKeyword = v
			}

		case "_prefix":
			if v, ok := termValue.(bool); ok {
				tv.IsPrefix = v
			}

		case "protected":
			if v, ok := termValue.(bool); ok {
				tv.IsProtected = v
			}

		case "@type":
			if v, ok := termValue.(string); ok {
				tv.TypeKeyword = v

			}

		//case "@container":
		//tv.ContainerKeyword:  //interface{}
		// TODO
		/*
			if v,ok := attr.(interface{}); ok {
				tv.ContainerKeyword = v
			}
		*/

		//case "@context":
		//tv.ContextKeyword:   interface{}
		// TODO
		/*
			if v,ok := attr.(interface{}); ok {
				tv.ContextKeyword = v
			}
		*/

		case "@language":

			if v, ok := termValue.(string); ok {
				tv.LanguageKeyword = v
			}

		case "@direction":

			if v, ok := termValue.(string); ok {
				tv.DirectionKeyword = v
			}

		case "@nest":

			if v, ok := termValue.(string); ok {
				tv.NestKeyword = v
			}

		case "@index":

			if v, ok := termValue.(string); ok {
				tv.IndexKeyword = v
			}

		}
	}
}
