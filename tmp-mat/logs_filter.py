import json

def filter_logs(log_file, filters):
    filtered_logs = []

    with open(log_file, "r", encoding="utf-8") as f:
        for line in f:
            log_entry = json.loads(line.strip())  # Читаем строку как JSON
            if any(value in log_entry.get(key, "") for key, value in filters.items()):
                continue  # Исключаем запись из выборки
            filtered_logs.append(log_entry)

    return filtered_logs

# Пример использования:
log_file = "/tmp/spqr/spqr.log"
filters = {
    "message": "etcd keep alive channel",
    "message": "etcdqdb: router already opened, nothing to do here",
    "message": "etcdqdb: open router",
    "message": "open router response",
    "message": "memqdb: update coordinator address",
    "message": "qdb coordinator: sync coordinator address",
}

filtered_logs = filter_logs(log_file, filters)

# Вывод отфильтрованных логов
for log in filtered_logs:
    print(json.dumps(log, ensure_ascii=False, indent=4))
