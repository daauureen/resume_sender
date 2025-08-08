# 🧪 Тестовое задание (Junior Golang Developer)

## 📌 Цель

Написать CLI-программу или HTTP-сервис(предпочтительнее) на Go, которая:

1. Принимает ссылку на ваше резюме
2. Генерирует хеш и уникальный ID
3. Создаёт JSON-отчёт
4. **Программно отправляет отчёт и исходный код на email: `szhaisan@wtotem.com`**

---

## 🔧 Условия

### 🧾 Аргументы командной строки:

- `--cv-url` — ссылка на ваше резюме, например:
  https://astana.hh.kz/resume/c34cbb24ff09a340310039ed1f457276644445
- `--email` — ваш email (будет указан в JSON и как отправитель письма)
- `--smtp-login` — логин вашей почты (например, Gmail)
- `--smtp-password` — пароль приложения (не основной пароль!)
- `--smtp-server` — SMTP-сервер (по умолчанию `smtp.gmail.com`)
- `--smtp-port` — порт (обычно `587`)

---

## ✅ Что программа должна делать

1. Вычислить **SHA256-хеш** от `cv-url`.
2. Сгенерировать `user_id`:
   - первые 8 символов хеша + 4 случайных символа
   - пример: `e3b0c442-a1b2`
3. Собрать JSON:

```json
{
    "cv_url": "https://...",
    "hash": "e3b0c44298fc1c149afbf4c8996fb...",
    "user_id": "e3b0c442-a1b2",
    "email": "ваш email",
    "timestamp": "2025-08-06T12:34:56Z"
}
```

4. Сохранить файл:
```report_<user_id>.json```
5. Отправить этот JSON-файл, и код проекта на email ``sabitov.olzhas@wtotem.com`` как прикреплённые файлы в письме.

## Пример Запуска

```bash
go run main.go \
  --cv-url=https://astana.hh.kz/resume/c34cbb24ff09a340310039ed1f457276644445 \
  --email=ivan@example.com \
  --smtp-login=ivan@gmail.com \
  --smtp-password=abcd1234xyz \
  --smtp-server=smtp.gmail.com \
  --smtp-port=587
```

## Отправка письма
- **Тема письма:** Golang Test – <user_id>  
- **Получатель:** szhaisan@wtotem.com  
- **В теле письма:** Автоматическая отправка отчёта  
- **Вложение:** report_<user_id>.json, source_code.zip

