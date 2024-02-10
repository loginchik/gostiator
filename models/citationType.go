package models

type CitationType struct {
	SystemName  string `json:"system_name,omitempty"`
	AppName     string `json:"app_name,omitempty"`
	WindowName  string `json:"window_name,omitempty"`
	Description string `json:"description"`
}

var ArticleCT = CitationType{
	SystemName:  "article",
	AppName:     "Статья",
	WindowName:  "Новая статья",
	Description: "Научная статья из рецензируемого журнала",
}

var BookCT = CitationType{
	SystemName:  "book",
	AppName:     "Книга",
	WindowName:  "Новая книга",
	Description: "Художественная, учебник, монография",
}

var ConferenceCT = CitationType{
	SystemName:  "conference",
	AppName:     "Материал конференции",
	WindowName:  "Новая книга",
	Description: "Статья, написанная как материал конференции",
}

var BookPartCT = CitationType{
	SystemName:  "bookPart",
	AppName:     "Часть книги",
	WindowName:  "Новая часть книги",
	Description: "Глава/раздел книги",
}

var WebsiteCT = CitationType{
	SystemName:  "website",
	AppName:     "Веб-сайт",
	WindowName:  "Новый веб-сайт",
	Description: "Веб-сайт как полноценный источник",
}

var WebtextCT = CitationType{
	SystemName:  "webContent",
	AppName:     "Интернет-текст",
	WindowName:  "Новый интернет-текст",
	Description: "Записи из соцсетей, материалы СМИ и похожее",
}

var WebvideoCT = CitationType{
	SystemName:  "webVideo",
	AppName:     "Интернет-видео",
	WindowName:  "Новое интернет-видео",
	Description: "Видео, опубликованное на хостинге",
}

var FilmCT = CitationType{
	SystemName:  "film",
	AppName:     "Фильм",
	WindowName:  "Новая фильм",
	Description: "Фильм, мультфильм",
}

var CitationTypeOptions = []CitationType{
	ArticleCT /*, ConferenceCT */, BookCT,
	BookPartCT, WebsiteCT, WebtextCT,
	WebvideoCT, FilmCT,
}
