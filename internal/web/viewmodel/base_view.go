package viewmodel

type BaseData struct {
	CSRFToken   string
	User        *UserView
	MenuItems   []MenuItemView
	CurrentLang string
	CurrentPath string
}

func NewBaseData(csrfToken string, user *UserView, menu []MenuItemView, lang string, path string) BaseData {
	return BaseData{
		CSRFToken:   csrfToken,
		User:        user,
		MenuItems:   menu,
		CurrentLang: lang,
		CurrentPath: path,
	}
}
