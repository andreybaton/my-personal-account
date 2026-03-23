todo:
) кокнурентную валидацию данных в регистрации
) рейт лимитинг добавить
) воркерпул для хеширования паролей
) кэширование токенов, юзеров, ролей и прав доступа, отозванных токенов
    type AuthCache struct {
    tokenCache *TokenCache локально (чаще всего) мапа с мьютексом
    userCache *UserRedisCache редис 
    permCache *PermissionsCache 
    rateLimitCache *RateLimitCache редис
    blacklistCache *BlacklistCache редис (синхронизация нужна)
}

