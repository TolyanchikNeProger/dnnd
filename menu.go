package dnnd

// MenuItem описывает элемент меню (например, "Главная", "Кампании", "Профиль")
type MenuItem struct {
	ID          int    `json:"id"`          // Уникальный идентификатор
	Title       string `json:"title"`       // Название пункта (например, "Мои игры")
	Path        string `json:"path"`        // URL-путь (например, "/games")
	Icon        string `json:"icon"`        // Иконка (название класса или ссылка)
	ParentID    int    `json:"parentId"`    // ID родительского пункта (для подменю)
	AccessLevel int    `json:"accessLevel"` // Уровень доступа (0=гость, 1=игрок, 2=ГМ)
}

// MainMenu структура для хранения всего меню приложения
type MainMenu struct {
	Items []MenuItem `json:"items"`
}

// Пример метода для фильтрации меню по уровню доступа
func (m *MainMenu) GetMenuForUser(accessLevel int) []MenuItem {
	var filtered []MenuItem
	for _, item := range m.Items {
		if item.AccessLevel <= accessLevel {
			filtered = append(filtered, item)
		}
	}
	return filtered
}
