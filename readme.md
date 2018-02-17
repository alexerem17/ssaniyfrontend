# Stacker Backend

# Todo
- [ ] Нужно выбрать header для аутификации
- [ ] Action.weight: func by SQL / views SQL / virtual index / trigger
```SQL
SELECT * FROM action WHERE user_id = id ORDER BY weight
```
- [ ] Нужно написать middleware для аутификации
- [x] Middleware for response JSON API recovery
- [ ] Подключить CRON GO Library
- [ ] Action.gps