# 🏆 Fantasy Bot

Telegram бот для получения результатов фэнтези-футбола с сайта Sports.ru. Бот предоставляет красивую и читаемую информацию о матчах, статистике игроков и рейтинге в сезоне.

## ✨ Возможности

- 📊 **Мониторинг матчей** - отслеживание результатов матчей в реальном времени
- ⚽ **Статистика игроков** - информация о набранных очках каждого игрока
- 🏆 **Рейтинг сезона** - текущая позиция в общем рейтинге
- 🎯 **Умный ранг** - автоматическое определение топ-1, топ-3, топ-5, топ-10
- 🌍 **Множественные лиги** - поддержка различных футбольных лиг
- 📱 **Красивое форматирование** - эмодзи и структурированный вывод

## 🚀 Поддерживаемые лиги

- 🇮🇹 **Италия** (Serie A)
- 🇷🇺 **Россия** (РПЛ)
- 🇬🇧 **Англия** (Premier League)
- 🇪🇸 **Испания** (La Liga)
- 🇫🇷 **Франция** (Ligue 1)
- 🇩🇪 **Германия** (Bundesliga)
- 🇵🇹 **Португалия** (Primeira Liga)
- 🇳🇱 **Голландия** (Eredivisie)
- 🇹🇷 **Турция** (Süper Lig)
- 🏆 **Лига чемпионов** (Champions League)

## 📋 Требования

- Go 1.19 или выше
- Telegram Bot Token
- Настроенные переменные окружения для каждой лиги

## 🛠 Установка

1. **Клонируйте репозиторий:**
```bash
git clone <repository-url>
cd FantasyBot
```

2. **Установите зависимости:**
```bash
go mod tidy
```

3. **Создайте файл `.env` с переменными окружения:**
```env
# Telegram Bot Token
TELEGRAM_BOT_API_TOKEN=your_bot_token_here

# Italy League
ITALY_ID=your_italy_squad_id
ITALY_TOURNAMENT=serie-a
ITALY_QUERY=your_italy_query

# Russia League
RUSSIA_ID=your_russia_squad_id
RUSSIA_TOURNAMENT=russian-premier-league
RUSSIA_QUERY=your_russia_query

# England League
ENGLAND_ID=your_england_squad_id
ENGLAND_TOURNAMENT=premier-league
ENGLAND_QUERY=your_england_query

# Spain League
SPAIN_ID=your_spain_squad_id
SPAIN_TOURNAMENT=la-liga
SPAIN_QUERY=your_spain_query

# France League
FRANCE_ID=your_france_squad_id
FRANCE_TOURNAMENT=ligue-1
FRANCE_QUERY=your_france_query

# Germany League
GERMANY_ID=your_germany_squad_id
GERMANY_TOURNAMENT=bundesliga
GERMANY_QUERY=your_germany_query

# Portugal League
PORTUGAL_ID=your_portugal_squad_id
PORTUGAL_TOURNAMENT=primeira-liga
PORTUGAL_QUERY=your_portugal_query

# Holland League
HOLLAND_ID=your_holland_squad_id
HOLLAND_TOURNAMENT=eredivisie
HOLLAND_QUERY=your_holland_query

# Turkey League
TURKEY_ID=your_turkey_squad_id
TURKEY_TOURNAMENT=super-lig
TURKEY_QUERY=your_turkey_query

# Championship League
CHAMPIONSHIP_ID=your_championship_squad_id
CHAMPIONSHIP_TOURNAMENT=champions-league
CHAMPIONSHIP_QUERY=your_championship_query
```

4. **Соберите проект:**
```bash
go build -o fantasybot
```

5. **Запустите бота:**
```bash
./fantasybot
```

## 📱 Использование

1. **Найдите вашего бота в Telegram** по имени
2. **Отправьте команду** `/points`
3. **Получите красивый отчет** с информацией о:
   - Текущем туре
   - Результатах матчей
   - Статистике ваших игроков
   - Позиции в рейтинге

## 📊 Пример вывода

```
🏆 **Italy**

📅 **Тур:** 3 тур. 🟢 Открыт
📊 **Статус:** 🟢 Тур открыт

⚽ **МАТЧИ:**
```
1. ✅ завершен Кремонезе 3 : 2 Сассуоло 
  --- Эмиль Аудеро (Кремонезе) scored **2** points. 🪑 On the bench. 🥅 Goalkeeper 
2. ✅ завершен Лечче 0 : 2 Милан 
  --- Сантьяго Хименес (Милан) scored **2** points. ⭐ Main cast. 🎯 Forward 
```

📈 **СТАТИСТИКА СЕЗОНА:**
🎯 Вы набрали 105 очков в сезоне и занимаете 4 219 место из 6 127
🏆 Топ-10: 65
```

## 🏗 Архитектура проекта

```
FantasyBot/
├── data.go          # Основная логика бота и обработка данных
├── formatter.go     # Форматирование сообщений для Telegram
├── helper.go        # Вспомогательные функции
├── scraper.go       # Парсинг данных с Sports.ru
├── season.go        # Обработка сезонной статистики
├── team.go          # Работа с командами и игроками
├── tour.go          # Информация о турах
├── go.mod           # Зависимости Go
├── go.sum           # Хеши зависимостей
└── README.md        # Документация
```

## 🔧 Основные функции

### `data.go`
- Основная логика Telegram бота
- Обработка команд пользователя
- Интеграция с Sports.ru API
- Отправка сообщений

### `formatter.go`
- Красивое форматирование сообщений
- Добавление эмодзи и структуры
- Логика определения рангов (топ-1, топ-3, топ-5, топ-10)
- Локализация на русский язык

### `scraper.go`
- Парсинг HTML страниц Sports.ru
- Извлечение информации о матчах
- Обработка статусов матчей
- Форматирование данных об игроках

### `season.go`
- Расчет сезонной статистики
- Форматирование чисел с пробелами
- Подготовка данных о рейтинге

## 🎯 Особенности

- **Умное определение ранга**: автоматически показывает "Топ-1", "Топ-3", "Топ-5", "Топ-10" в зависимости от процента
- **Красивое форматирование**: использование эмодзи и структурированного вывода
- **Поддержка множественных лиг**: один бот для всех ваших фэнтези-команд
- **Реальное время**: актуальная информация о матчах и статистике

## 🤝 Вклад в проект

1. Форкните репозиторий
2. Создайте ветку для новой функции (`git checkout -b feature/amazing-feature`)
3. Зафиксируйте изменения (`git commit -m 'Add amazing feature'`)
4. Отправьте в ветку (`git push origin feature/amazing-feature`)
5. Откройте Pull Request

## 📄 Лицензия

Этот проект распространяется под лицензией MIT. См. файл `LICENSE` для получения дополнительной информации.

## 🆘 Поддержка

Если у вас возникли вопросы или проблемы:

1. Проверьте раздел [Issues](../../issues)
2. Создайте новый issue с подробным описанием проблемы
3. Убедитесь, что все переменные окружения настроены правильно

---

**Создано с ❤️ для фэнтези-футбола**