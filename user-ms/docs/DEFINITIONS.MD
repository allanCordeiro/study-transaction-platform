# User's management microservice definitions
## Objetivo
Descrever de forma simples estrutura e regras de negócio do microserviço de gestão de usuário

## Usuário
Via de regra existem dois perfis de usuário. 
- Cliente: usuário comum da plataforma;
- Distribuidor: usuário com característica de lojista.

### Estrutura esperada
- id: UUID
- name: string
- email: string
- type: string (enum)
- password: string (hash)
- created_at: date/time
- updated_at: date/time
- deleted_at: date/time
- is_active: boolean

### Operações
#### Server
O server estará habilitado para cumprir as seguintes operações:
 - Alterar dados do usuário (exceto seu tipo e email, que é chave primario)
 - Realizar autenticação

Todo o processo será disponibilizado via API HTTP

### CLI
Já o CLI será um binário que trará a possibilidade de:
 - Criar usuário
 - Alterar o perfil do usuário
 - Desativar ou reativar usuário
