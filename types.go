package alice

// APIRequest определяет формат запроса
type APIRequest struct {
	Meta    Meta    `json:"meta"`
	Request Request `json:"request"`
	Session Session `json:"session"`
	Version string  `json:"version"`
}

// APIResponse определяет формат ответа
type APIResponse struct {
	Response Response `json:"response"`
	Session  Session  `json:"session"`
	Version  string   `json:"version"`
}

// Meta описывает информация об устройстве пользователя
type Meta struct {
	Locale   string `json:"locale"`
	Timezone string `json:"timezone"`
	ClientID string `json:"client_id"`
}

// Request описывает данные, полученные от пользователя
type Request struct {
	Command           string  `json:"command"`
	OriginalUtterance string  `json:"original_utterance"`
	Type              string  `json:"type"`
	Markup            Markup  `json:"markup"`
	Payload           Payload `json:"payload"`
	NLU               NLU     `json:"nlu"`
}

// Session описывает данные сессии пользователя
type Session struct {
	New       bool   `json:"new"`
	MessageID int64  `json:"message_id"`
	SessionID string `json:"session_id"`
	SkillID   string `json:"skill_id"`
	UserID    string `json:"user_id"`
}

// Response описывает данные для ответа пользователю
type Response struct {
	Card       *Card    `json:"card,omitempty"`
	Text       string   `json:"text,omitempty"`
	TTS        string   `json:"tts,omitempty"`
	Buttons    []Button `json:"buttons,omitempty"`
	EndSession bool     `json:"end_session,omitempty"`
}

// Markup определяет характеристики, выявленные Яндекс.Диалогами
type Markup struct {
	DangerousContext bool `json:"dangerous_context"`
}

// NLU определяет именованные сущности
type NLU struct {
	Tokens   []string  `json:"tokens"`
	Entities []Entitie `json:"entities"`
}

// Entitie определяет массив именованных сущностей
type Entitie struct {
	Tokens   Token  `json:"tokens"`
	Type     string `json:"type"`
	ValueInt int    `json:"value"`
	Value    Value  `json:"value"`
}

// Token определяет начало и конец именованной сущности в массиве слов
type Token struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// Value формально описывает именованную сущность
type Value struct {
	HouseNumber     string `json:"house_number"`
	Street          string `json:"street"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Day             int    `json:"day"`
	DayIsRelative   bool   `json:"day_is_relative"`
	Month           int    `json:"month"`
	MonthIsRelative bool   `json:"month_is_relative"`
}

// Card описывает карточку — сообщения с поддержкой изображений
type Card struct {
	Type        string     `json:"type,omitempty"`
	ImageID     string     `json:"image_id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Button      CardButton `json:"button,omitempty"`
	Header      Header     `json:"header,omitempty"`
	Items       []Item     `json:"items,omitempty"`
	Footer      Footer     `json:"footer,omitempty"`
}

// Button описывает свойства для кнопки
type Button struct {
	Title   string  `json:"title,omitempty"`
	Payload Payload `json:"payload,omitempty"`
	URL     string  `json:"url,omitempty"`
	Hide    bool    `json:"hide,omitempty"`
}

// CardButton описывает свойство для кнопки под карточку
type CardButton struct {
	Text    string  `json:"text,omitempty"`
	URL     string  `json:"url,omitempty"`
	Payload Payload `json:"payload,omitempty"`
}

// Header описывает заголовок галереи изображений
type Header struct {
	Text string `json:"text,omitempty"`
}

// Item описывает карточку для набора изображений
type Item struct {
	ImageID     string     `json:"image_id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Button      CardButton `json:"button,omitempty"`
}

// Footer описывает кнопки под галереей изображений
type Footer struct {
	Text   string     `json:"text,omitempty"`
	Button CardButton `json:"button,omitempty"`
}

// Payload определяет JSON, полученный с нажатой кнопкой
// от обработчика навыка (в ответе на предыдущий запрос), максимум 4096 байт
type Payload interface{}
