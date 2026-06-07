# S-Store Admin

Admin dashboard VueJS theo Clean Architecture.

## Cấu trúc

```
src/
├── app/          # Khởi động app, router, providers, config
├── core/         # Constants, errors, types, utils
├── domain/       # Entities, repository interfaces, use cases
├── data/         # API, models, mappers, repository implementations
├── presentation/ # Pages, components, layouts, stores, composables
└── shared/       # Styles, plugins, assets
```

## Bắt đầu

```bash
npm install
npm run dev
```

## Scripts

| Script | Mô tả |
|---|---|
| `npm run dev` | Chạy dev server |
| `npm run build` | Build production |
| `npm run type-check` | Kiểm tra TypeScript |
| `npm run lint` | ESLint |
| `npm run format` | Prettier |

## Env

- `.env.development` — môi trường dev
- `.env.production` — môi trường production

## Luồng auth

```
LoginPage → AuthStore → LoginUseCase → UserRepository → AuthApi → Axios
```

Xem thêm: [docs/vue_clean_architecture_setup.md](./docs/vue_clean_architecture_setup.md)
