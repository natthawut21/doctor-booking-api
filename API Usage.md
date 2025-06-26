
# 1. Get all doctors ( แนะนำใช้ Grails )

**GET**

```sh
curl --location 'http://localhost:8181/doctor'
```

*Result*

```json
[
    {
        "id": 4,
        "version": 1,
        "name": "Dr. POP",
        "specialty": "อยุกรรม",
        "subSpecialty": "",
        "department": "",
        "phone": "081-xxx-xxxx",
        "email": "dr-pop-vnurser-1@gmail.com",
        "bankAccountName": "ธนาคารกสิกรไทย",
        "bankAccountNumber": "01-123-456789",
        "licenseNumber": "",
        "licenseIssuer": "แพทยสภา",
        "licenseIssuedDate": "2022-02-01T00:00:00Z",
        "licenseExpiryDate": "2027-12-01T00:00:00Z"
    },
    {
        "id": 14,
        "version": 3,
        "name": "Dr. POP",
        "specialty": "Internal Medicine",
        "subSpecialty": "Arrhythmia",
        "department": "Cardiology",
        "phone": "081-4444444",
        "email": "pop.dr@gamil.com",
        "bankAccountName": "ธนาคารกสิกรไทย",
        "bankAccountNumber": "01-123-456789",
        "licenseNumber": "",
        "licenseIssuer": "แพทยสภา",
        "licenseIssuedDate": "2022-02-01T00:00:00Z",
        "licenseExpiryDate": "2027-12-01T00:00:00Z"
    }
]
```

# 2.  Doctor Schedule ( แนะนำใช้ Grails )
##  2.1 Create Doctor Schedule ( แนะนำใช้ Grails )

**POST**

```sh
curl --location 'http://localhost:8181/schedules' \
--header 'Content-Type: application/json' \
--data '{
    "doctor_id": 14,
    "day_of_week": "FRIDAY",
    "start_time": "09:00",
    "end_time": "12:00"
}'
```


*Result*
```json
{
    "id": 29,
    "version": 0,
    "doctor_id": 14,
    "day_of_week": "FRIDAY",
    "start_time": "09:00",
    "end_time": "12:00"
}

```

##  2.2 Get Doctor Schedule 
**GET**
```bash
curl --location 'http://localhost:8181/doctor/{doctorId}/schedules'
```

*Result*
```json
[
    {
        "id": 17,
        "version": 0,
        "doctor_id": 14,
        "day_of_week": "MONDAY",
        "start_time": "10:00:00",
        "end_time": "12:00:00"
    },
    {
        "id": 18,
        "version": 0,
        "doctor_id": 14,
        "day_of_week": "MONDAY",
        "start_time": "13:00:00",
        "end_time": "17:00:00"
    },
    {
        "id": 26,
        "version": 0,
        "doctor_id": 14,
        "day_of_week": "TUESDAY",
        "start_time": "10:00:00",
        "end_time": "12:00:00"
    },
    {
        "id": 20,
        "version": 1,
        "doctor_id": 14,
        "day_of_week": "TUESDAY",
        "start_time": "13:00:00",
        "end_time": "14:00:00"
    },
    {
        "id": 21,
        "version": 0,
        "doctor_id": 14,
        "day_of_week": "WEDNESDAY",
        "start_time": "09:00:00",
        "end_time": "12:00:00"
    },
    {
        "id": 22,
        "version": 0,
        "doctor_id": 14,
        "day_of_week": "WEDNESDAY",
        "start_time": "13:00:00",
        "end_time": "17:00:00"
    },
    {
        "id": 23,
        "version": 0,
        "doctor_id": 14,
        "day_of_week": "THURSDAY",
        "start_time": "09:00:00",
        "end_time": "12:00:00"
    },
    {
        "id": 24,
        "version": 0,
        "doctor_id": 14,
        "day_of_week": "THURSDAY",
        "start_time": "13:00:00",
        "end_time": "18:00:00"
    },
    {
        "id": 29,
        "version": 0,
        "doctor_id": 14,
        "day_of_week": "FRIDAY",
        "start_time": "09:00:00",
        "end_time": "12:00:00"
    },
    {
        "id": 25,
        "version": 0,
        "doctor_id": 14,
        "day_of_week": "FRIDAY",
        "start_time": "13:00:00",
        "end_time": "19:00:00"
    }
]
```



# 3. Get Doctor Slot

**GET**

```sh
curl --location 'http://localhost:8181/slots/all?doctorId={:doctorId}&date=2025-06-30'
```

*Result*

```json
[
    {
        "id": 1664,
        "doctor_id": 4,
        "start_time": "2025-06-30T09:00:00+07:00",
        "end_time": "2025-06-30T09:15:00+07:00",
        "status": "CONFIRMED",
        "changed_at": "2025-06-26T16:38:33.859Z"
    },
    {
        "id": 1665,
        "doctor_id": 4,
        "start_time": "2025-06-30T09:20:00+07:00",
        "end_time": "2025-06-30T09:35:00+07:00",
        "status": "AVAILABLE"
    },
    {
        "id": 1666,
        "doctor_id": 4,
        "start_time": "2025-06-30T09:40:00+07:00",
        "end_time": "2025-06-30T09:55:00+07:00",
        "status": "AVAILABLE"
    },
    {
        "id": 1667,
        "doctor_id": 4,
        "start_time": "2025-06-30T10:00:00+07:00",
        "end_time": "2025-06-30T10:15:00+07:00",
        "status": "AVAILABLE"
    },
    {
        "id": 1668,
        "doctor_id": 4,
        "start_time": "2025-06-30T10:20:00+07:00",
        "end_time": "2025-06-30T10:35:00+07:00",
        "status": "AVAILABLE"
    },
    {
        "id": 1669,
        "doctor_id": 4,
        "start_time": "2025-06-30T10:40:00+07:00",
        "end_time": "2025-06-30T10:55:00+07:00",
        "status": "AVAILABLE"
    },
    {
        "id": 1670,
        "doctor_id": 4,
        "start_time": "2025-06-30T11:00:00+07:00",
        "end_time": "2025-06-30T11:15:00+07:00",
        "status": "AVAILABLE"
    },
    {
        "id": 1671,
        "doctor_id": 4,
        "start_time": "2025-06-30T11:20:00+07:00",
        "end_time": "2025-06-30T11:35:00+07:00",
        "status": "AVAILABLE"
    },
    {
        "id": 1672,
        "doctor_id": 4,
        "start_time": "2025-06-30T11:40:00+07:00",
        "end_time": "2025-06-30T11:55:00+07:00",
        "status": "AVAILABLE"
    },
    {
        "id": 1673,
        "doctor_id": 4,
        "start_time": "2025-06-30T13:00:00+07:00",
        "end_time": "2025-06-30T13:15:00+07:00",
        "status": "AVAILABLE"
    },
    {
        "id": 1674,
        "doctor_id": 4,
        "start_time": "2025-06-30T13:20:00+07:00",
        "end_time": "2025-06-30T13:35:00+07:00",
        "status": "AVAILABLE"
    },
    {
        "id": 1675,
        "doctor_id": 4,
        "start_time": "2025-06-30T13:40:00+07:00",
        "end_time": "2025-06-30T13:55:00+07:00",
        "status": "AVAILABLE"
    },
    {
        "id": 1676,
        "doctor_id": 4,
        "start_time": "2025-06-30T14:00:00+07:00",
        "end_time": "2025-06-30T14:15:00+07:00",
        "status": "AVAILABLE"
    },
    {
        "id": 1677,
        "doctor_id": 4,
        "start_time": "2025-06-30T14:20:00+07:00",
        "end_time": "2025-06-30T14:35:00+07:00",
        "status": "AVAILABLE"
    },
    {
        "id": 1678,
        "doctor_id": 4,
        "start_time": "2025-06-30T14:40:00+07:00",
        "end_time": "2025-06-30T14:55:00+07:00",
        "status": "AVAILABLE"
    },
    {
        "id": 1679,
        "doctor_id": 4,
        "start_time": "2025-06-30T15:00:00+07:00",
        "end_time": "2025-06-30T15:15:00+07:00",
        "status": "CONFIRMED"
    },
    {
        "id": 1680,
        "doctor_id": 4,
        "start_time": "2025-06-30T15:20:00+07:00",
        "end_time": "2025-06-30T15:35:00+07:00",
        "status": "AVAILABLE"
    },
    {
        "id": 1681,
        "doctor_id": 4,
        "start_time": "2025-06-30T15:40:00+07:00",
        "end_time": "2025-06-30T15:55:00+07:00",
        "status": "AVAILABLE"
    },
    {
        "id": 1682,
        "doctor_id": 4,
        "start_time": "2025-06-30T16:00:00+07:00",
        "end_time": "2025-06-30T16:15:00+07:00",
        "status": "AVAILABLE"
    },
    {
        "id": 1683,
        "doctor_id": 4,
        "start_time": "2025-06-30T16:20:00+07:00",
        "end_time": "2025-06-30T16:35:00+07:00",
        "status": "AVAILABLE"
    },
    {
        "id": 1684,
        "doctor_id": 4,
        "start_time": "2025-06-30T16:40:00+07:00",
        "end_time": "2025-06-30T16:55:00+07:00",
        "status": "AVAILABLE"
    }
]
```

# 4.  Appointment
การทำนัดหรือการจองเวลาของหมอมีอยู่ 2 วิธี
## 4.1 POST Appointment

**POST**
```sh
curl --location 'http://localhost:8181/appointments/book' \
--header 'Content-Type: application/json' \
--data '{
   "slotId": {SlotId},
   "doctorId": {DoctorId},
    "username": "{Username}/{Patient}"
}
'
```

*Result*

```json
{
    "id": 9,
    "message": "Appointment booked"
}
```
`status` ของ Slot จะเปลี่ยนเป็น `CONFIRMED`
## 4.2 PUT Update Slot

วิธีที่ 2 ในการ Update สถานะของ Slot ของหมอ จะใช้การ Update ไปที่ Slot id โดย กำหนดค่าให้กับ `status` ด้วยค่าตัวแปรตามรายการ
	
	`AVAILABLE`
	`PENDING`
	`CONFIRMED`
	`CANCELED`


**PUT**

```sh
curl --location --request PUT 'http://localhost:8181/slots/{SlotId}/status' \
--header 'Content-Type: application/json' \
--data '{
  "status": "CONFIRMED"
}'
```

*Result*
```json
{
    "message": "Slot status updated"
}
```

## 4.3 Get Slot Info and History
**GET**
```sh
curl --location 'http://localhost:8181/slots/{SlotId}/info'
```

*Result*
```Json
{
    "id": 1664,
    "doctor_id": 4,
    "start_time": "2025-06-30T09:00:00+07:00",
    "end_time": "2025-06-30T09:15:00+07:00",
    "status": "CONFIRMED",
    "history": [
        {
            "ID": 6,
            "slot_id": 1664,
            "old_status": "AVAILABLE",
            "new_status": "CONFIRMED",
            "changed_by": "",
            "changed_at": "2025-06-26T23:38:33.859+07:00",
            "Slot": {
                "id": 0,
                "version": 0,
                "doctor_id": 0,
                "start_time": "0001-01-01T00:00:00Z",
                "end_time": "0001-01-01T00:00:00Z",
                "status": ""
            }
        }
    ]
}
```

# 5. List all Appointment

```sh
curl --location 'http://localhost:8181/appointments' \
--header 'Content-Type: application/json'
```

```json
[
    {
        "id": 4,
        "version": 0,
        "doctor_id": 4,
        "patient_id": 3,
        "slot_id": 1643,
        "created_at": "2025-06-23T20:20:07.582211Z",
        "Doctor": {
            "id": 4,
            "version": 1,
            "name": "Dr. POP",
            "specialty": "อยุกรรม",
            "subSpecialty": "",
            "department": "",
            "phone": "081-xxx-xxxx",
            "email": "dr-pop-vnurser-1@gmail.com",
            "bankAccountName": "ธนาคารกสิกรไทย",
            "bankAccountNumber": "01-123-456789",
            "licenseNumber": "",
            "licenseIssuer": "แพทยสภา",
            "licenseIssuedDate": "2022-02-01T00:00:00Z",
            "licenseExpiryDate": "2027-12-01T00:00:00Z"
        },
        "Patient": {
            "id": 3,
            "username": "testpatient",
            "name": ""
        },
        "Slot": {
            "id": 1643,
            "version": 0,
            "doctor_id": 4,
            "start_time": "2025-06-26T06:00:00Z",
            "end_time": "2025-06-26T06:15:00Z",
            "status": "PENDING"
        }
    },
    {
        "id": 5,
        "version": 0,
        "doctor_id": 4,
        "patient_id": 3,
        "slot_id": 1665,
        "created_at": "2025-06-26T16:23:12.485832Z",
        "Doctor": {
            "id": 4,
            "version": 1,
            "name": "Dr. POP",
            "specialty": "อยุกรรม",
            "subSpecialty": "",
            "department": "",
            "phone": "081-xxx-xxxx",
            "email": "dr-pop-vnurser-1@gmail.com",
            "bankAccountName": "ธนาคารกสิกรไทย",
            "bankAccountNumber": "01-123-456789",
            "licenseNumber": "",
            "licenseIssuer": "แพทยสภา",
            "licenseIssuedDate": "2022-02-01T00:00:00Z",
            "licenseExpiryDate": "2027-12-01T00:00:00Z"
        },
        "Patient": {
            "id": 3,
            "username": "testpatient",
            "name": ""
        },
        "Slot": {
            "id": 1665,
            "version": 0,
            "doctor_id": 4,
            "start_time": "2025-06-30T02:20:00Z",
            "end_time": "2025-06-30T02:35:00Z",
            "status": "AVAILABLE"
        }
    },
    {
        "id": 6,
        "version": 0,
        "doctor_id": 4,
        "patient_id": 3,
        "slot_id": 1679,
        "created_at": "2025-06-26T16:29:17.378606Z",
        "Doctor": {
            "id": 4,
            "version": 1,
            "name": "Dr. POP",
            "specialty": "อยุกรรม",
            "subSpecialty": "",
            "department": "",
            "phone": "081-xxx-xxxx",
            "email": "dr-pop-vnurser-1@gmail.com",
            "bankAccountName": "ธนาคารกสิกรไทย",
            "bankAccountNumber": "01-123-456789",
            "licenseNumber": "",
            "licenseIssuer": "แพทยสภา",
            "licenseIssuedDate": "2022-02-01T00:00:00Z",
            "licenseExpiryDate": "2027-12-01T00:00:00Z"
        },
        "Patient": {
            "id": 3,
            "username": "testpatient",
            "name": ""
        },
        "Slot": {
            "id": 1679,
            "version": 0,
            "doctor_id": 4,
            "start_time": "2025-06-30T08:00:00Z",
            "end_time": "2025-06-30T08:15:00Z",
            "status": "CONFIRMED"
        }
    }
]
```