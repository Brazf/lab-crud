## Resumo da Arquitetura em Camadas

| Camada         | Responsabilidade              | Depende de        |
| -------------- | ----------------------------- | ----------------- |
| Config         | Configurações básicas         | —                 |
| Infrastructure | Logs, middlewares utilitários | —                 |
| Database       | Conexão com MySQL             | Config            |
| Model          | Entidades                     | —                 |
| Repository     | Acesso ao banco               | Model, Database   |
| Service        | Regras de negócio             | Repository, Model |
| Handler        | HTTP/REST                     | Service           |
| Main           | Monta a aplicação             | Todas acima       |


