# สำหรับการ export keycloak

# สำหรับการ export keycloak

## คำสั่ง Export:

### Export realm เดียว (Devpool_project):
```bash
docker exec keycloak /opt/keycloak/bin/kc.sh export --dir /tmp/realm-export --realm Devpool_project --users realm_file
```

### Export ทุก realms:
```bash
docker exec keycloak /opt/keycloak/bin/kc.sh export --dir /tmp/realm-export --users realm_file
```

## Copy ไฟล์ออกมา:

### สำหรับ realm เดียว:
```bash
docker cp keycloak:/tmp/realm-export/Devpool_project-realm.json ./Keycloak/realm-export/
```

### สำหรับทุก realms:
```bash
# ตรวจสอบไฟล์ที่มี
docker exec keycloak ls -la /tmp/export-all/

# Copy ทุกไฟล์
docker cp keycloak:/tmp/export-all/ ./Keycloak/realm-export-all/
```

## การจัดการไฟล์สำหรับ Auto Import:

### ❌ Export ทุก realms - ไม่ auto import
เมื่อ export ทุก realms จะได้ไฟล์หลายไฟล์:
- `master-realm.json` (realm ของ Keycloak เอง)
- `Devpool_project-realm.json` (realm ของเรา)
- อื่นๆ (ถ้ามี)

**ปัญหา**: Keycloak จะ import ทุกไฟล์ ซึ่งอาจทำให้เกิดปัญหา

### ✅ Export realm เดียว - auto import ได้
```bash
# Copy เฉพาะ realm ที่ต้องการเป็น realm-export.json
cp Devpool_project-realm.json realm-export.json
```

## แนะนำ:
1. **สำหรับ production**: ใช้ export realm เดียว
2. **สำหรับ backup**: ใช้ export ทุก realms
3. **Auto import**: ใช้เฉพาะ realm ที่ต้องการ

## ไฟล์ที่ได้:
- `Devpool_project-realm.json` - ไฟล์ realm ต้นฉบับ (เก็บไว้เป็น backup)
- `realm-export.json` - ไฟล์สำหรับ auto import (copy จาก Devpool_project-realm.json)

## การใช้งาน:
1. **เก็บ backup**: ไฟล์ `Devpool_project-realm.json` 
2. **Auto import**: ไฟล์ `realm-export.json` จะถูก import อัตโนมัติเมื่อ `docker-compose up`
3. **ไม่ต้องเปลี่ยนชื่อ**: ระบบจัดการให้อัตโนมัติ

## Volume Mapping:
```yaml
volumes:
  - ./Keycloak/realm-export:/opt/keycloak/data/import
```

✅ พร้อมใช้งาน - ไม่ต้องเปลี่ยนชื่อไฟล์