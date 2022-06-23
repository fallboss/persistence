# Boss DB persistence

Database persistence tool.

## Environment Variables

### Database Variables
#### Database Connection
| Variable | Mandatory | Description | Default |
| --- | :---: | --- | :---: |
| DB_HOST | :white_check_mark: | Host of the database | - |
| DB_PORT | :white_check_mark: | Port of the database | - |
| DB_USER | :white_check_mark: | User of the database | - |
| DB_PASSWORD | :white_check_mark: | Password of the database | - |
| DB_NAME | :white_check_mark: | Name of the database | - |
| DB_SSL_MODE | :no_entry: | SSL mode of the database | require |

#### Database Performance
| Variable | Mandatory | Description | Default |
| --- | :---: | --- | :---: |
| DB_SELECT_LIMIT | :no_entry: | Limit of each select query | 10000 |
| DB_MAX_CONNS | :no_entry: | Maximum number of connections to the database | 30 |
| DB_WAIT_AFTER_QUERY | :no_entry: | Time to block each connection after a query in milliseconds | 10 |

###Example Environment Variables
- name: DB_MAX_CONNS
  value: "40"
- name: DB_WAIT_AFTER_QUERY
  value: "300"
- name: DB_SELECT_LIMIT
  value: "10000"
- name: DB_HOST
  valueFrom:
    secretKeyRef:
      key: latest
      name: sm-database-host-cl
- name: DB_PORT
  value: "5432"
- name: DB_USER
  valueFrom:
    secretKeyRef:
      key: latest
      name: sm-database-user-cl
- name: DB_PASSWORD
  valueFrom:
    secretKeyRef:
      key: latest
      name: sm-database-password-cl
- name: DB_NAME
  valueFrom:
    secretKeyRef:
      key: latest
      name: sm-database-scheme-cl