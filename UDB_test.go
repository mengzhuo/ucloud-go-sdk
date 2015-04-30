package ucloud

import (
	"fmt"
	"testing"
)

// ---------------- TestBackupUDBInstance ------------------
func TestBackupUDBInstance(t *testing.T) {
	fmt.Println("UDB......")
	r := &BackupUDBInstance{Region: "cn-east-01",
		BackupName: "master-backup",
		DBId:       "00f9868c-c7f5-4852-9eac-d200b678f0e1",
	}
	cmp := `https://api.ucloud.cn/?Action=BackupUDBInstance&Region=cn-east-01&DBId=00f9868c-c7f5-4852-9eac-d200b678f0e1&BackupName=master-backup`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestClearUDBLog ------------------
func TestClearUDBLog(t *testing.T) {
	r := &ClearUDBLog{Region: "cn-east-01",
		LogType: 30,
		DBId:    "00f9868c-c7f5-4852-9eac-d200b678f0e1",
	}
	cmp := `https://api.ucloud.cn/?Action=ClearUDBLog&Region=cn-east-01&DBId=00f9868c-c7f5-4852-9eac-d200b678f0e1&LogType=30`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestCreateUDBInstance ------------------
func TestCreateUDBInstance(t *testing.T) {
	r := &CreateUDBInstance{AdminUser: "root",
		DiskSpace:     500,
		AdminPassword: "mysql_12",
		Region:        "cn-east-01",
		ParamGroupId:  3,
		MemoryLimit:   600,
		ChargeType:    "Month",
		Name:          "udb-xxxxxxx",
		Port:          3306,
		DBTypeId:      "mysql-5.1",
	}
	cmp := `https://api.ucloud.cn/?Action=CreateUDBInstance&Region=cn-east-01&DBTypeId=mysql-5.1&ChargeType=Month&Name=udb-xxxxxxx&AdminUser=root&AdminPassword=mysql_12&Port=3306&ParamGroupId=3&MemoryLimit=600&DiskSpace=500`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestCreateUDBParamGroup ------------------
func TestCreateUDBParamGroup(t *testing.T) {
	r := &CreateUDBParamGroup{
		Region:      "cn-north-01",
		Description: "xxx-config",
		DBTypeId:    "mysql-5.1",
		SrcGroupId:  2,
		GroupName:   "mysql-xxx-config",
	}
	cmp := `https://api.ucloud.cn/?Region=cn-north-01&Action=CreateUDBParamGroup&DBTypeId=mysql-5.1&SrcGroupId=2&GroupName=mysql-xxx-config&Description=xxx-config`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestCreateUDBReplicationInstance ------------------
func TestCreateUDBReplicationInstance(t *testing.T) {
	r := &CreateUDBReplicationInstance{
		SrcId:     "1e06ad18-10cb-481c-89d9-e3ea64324fa3",
		Region:    "cn-east-01",
		IsArbiter: true,
		Name:      "mongodb-xxxxxxx-01",
	}
	cmp := `https://api.spark.ucloud.cn/?Action=CreateUDBReplicationInstance&Region=cn-east-01&SrcId=1e06ad18-10cb-481c-89d9-e3ea64324fa3&Name=mongodb-xxxxxxx-01`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestCreateUDBSlave ------------------
func TestCreateUDBSlave(t *testing.T) {
	r := &CreateUDBSlave{Region: "cn-east-01",
		SrcId: "00f9868c-c7f5-4852-9eac-d200b678f0e1",
		Name:  "udb-xxxxxxx-slave",
		Port:  3306,
	}
	cmp := `https://api.ucloud.cn/?Action=CreateUDBSlave&Region=cn-east-01&SrcId=00f9868c-c7f5-4852-9eac-d200b678f0e1&Name=udb-xxxxxxx-slave&Port=3306`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDeleteUDBBackup ------------------
func TestDeleteUDBBackup(t *testing.T) {
	r := &DeleteUDBBackup{Region: "cn-east-01",
		BackupId: 5160,
	}
	cmp := `https://api.ucloud.cn/?Action=DeleteUDBBackup&Region=cn-east-01&BackupId=5160`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDeleteUDBInstance ------------------
func TestDeleteUDBInstance(t *testing.T) {
	r := &DeleteUDBInstance{Region: "cn-east-01",
		DBId: "00f9868c-c7f5-4852-9eac-d200b678f0e1",
	}
	cmp := `https://api.ucloud.cn/?Action=DeleteUDBInstance&Region=cn-east-01&DBId=00f9868c-c7f5-4852-9eac-d200b678f0e1`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDeleteUDBParamGroup ------------------
func TestDeleteUDBParamGroup(t *testing.T) {
	r := &DeleteUDBParamGroup{Region: "cn-east-01",
		GroupId: 192,
	}
	cmp := `https://api.ucloud.cn/?Action=DeleteUDBParamGroup&Region=cn-east-01&GroupId=192`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeUDBBackupBlacklist ------------------
func TestDescribeUDBBackupBlacklist(t *testing.T) {
	r := &DescribeUDBBackupBlacklist{Region: "cn-east-01",
		DBId: "00f9868c-c7f5-4852-9eac-d200b678f0e1",
	}
	cmp := `https://api.ucloud.cn/?Action=DescribeUDBBackupBlacklist&Region=cn-east-01&DBId=00f9868c-c7f5-4852-9eac-d200b678f0e1`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeUDBBackup ------------------
func TestDescribeUDBBackup(t *testing.T) {
	r := &DescribeUDBBackup{Region: "cn-east-01",
		Limit:  20,
		Offset: 0,
	}
	cmp := `https://api.ucloud.cn/?Action=DescribeUDBBackup&Region=cn-east-01&Offset=0&Limit=20`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeUDBInstancePrice ------------------
func TestDescribeUDBInstancePrice(t *testing.T) {
	r := &DescribeUDBInstancePrice{Count: 1,
		DiskSpace:   20,
		Region:      "cn-east-01",
		MemoryLimit: 600,
		ChargeType:  "Month",
		Quantity:    12,
	}
	cmp := `https://api.ucloud.cn/?Action=DescribeUDBInstancePrice&Region=cn-east-01&Count=1&ChargeType=Month&Quantity=12&MemoryLimit=600&DiskSpace=20`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeUDBInstance ------------------
func TestDescribeUDBInstance(t *testing.T) {
	r := &DescribeUDBInstance{Region: "cn-east-01",
		DBId: "00f9868c-c7f5-4852-9eac-d200b678f0e1",
	}
	cmp := `https://api.ucloud.cn/?Action=DescribeUDBInstance&Region=cn-east-01&DBId=00f9868c-c7f5-4852-9eac-d200b678f0e1`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeUDBInstanceState ------------------
func TestDescribeUDBInstanceState(t *testing.T) {
	r := &DescribeUDBInstanceState{Region: "cn-east-01",
		DBId: "00f9868c-c7f5-4852-9eac-d200b678f0e1",
	}
	cmp := `https://api.ucloud.cn/?Action=DescribeUDBInstanceState&Region=cn-east-01&DBId=00f9868c-c7f5-4852-9eac-d200b678f0e1`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeUDBParamGroup ------------------
func TestDescribeUDBParamGroup(t *testing.T) {
	r := &DescribeUDBParamGroup{Region: "cn-east-01",
		GroupId: 20,
	}
	cmp := `https://api.ucloud.cn/?Action=DescribeUDBParamGroup&Region=cn-east-01&GroupId=20`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestDescribeUDBType ------------------
func TestDescribeUDBType(t *testing.T) {
	r := &DescribeUDBType{Region: "cn-east-01"}
	cmp := `https://api.ucloud.cn/?Action=DescribeUDBType&Region=cn-east-01`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestEditUDBBackupBlacklist ------------------
func TestEditUDBBackupBlacklist(t *testing.T) {
	r := &EditUDBBackupBlacklist{Blacklist: "abc.%",
		Region: "cn-east-01",
		DBId:   "00f9868c-c7f5-4852-9eac-d200b678f0e1",
	}
	cmp := `https://api.ucloud.cn/?Action=EditUDBBackupBlacklist&Region=cn-east-01&DBId=00f9868c-c7f5-4852-9eac-d200b678f0e1&Blacklist=abc.%;user.%;id.%;`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestModifyUDBInstanceName ------------------
func TestModifyUDBInstanceName(t *testing.T) {
	r := &ModifyUDBInstanceName{Region: "cn-east-01",
		Name: "xxx-yyy",
		DBId: "00f9868c-c7f5-4852-9eac-d200b678f0e1",
	}
	cmp := `https://api.ucloud.cn/?Action=ModifyUDBInstanceName&Region=cn-east-01&DBId=00f9868c-c7f5-4852-9eac-d200b678f0e1&Name=xxx-yyy`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestPromoteUDBSlave ------------------
func TestPromoteUDBSlave(t *testing.T) {
	r := &PromoteUDBSlave{Region: "cn-east-01",
		IsForce: true,
		DBId:    "00f9868c-c7f5-4852-9eac-d200b678f0e1",
	}
	cmp := `https://api.ucloud.cn/?Action=PromoteUDBSlave&Region=cn-east-01&DBId=00f9868c-c7f5-4852-9eac-d200b678f0e1&IsForce=true`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestRestartUDBInstance ------------------
func TestRestartUDBInstance(t *testing.T) {
	r := &RestartUDBInstance{Region: "cn-east-01",
		DBId: "00f9868c-c7f5-4852-9eac-d200b678f0e1",
	}
	cmp := `https://api.ucloud.cn/?Action=RestartUDBInstance&Region=cn-east-01&DBId=00f9868c-c7f5-4852-9eac-d200b678f0e1`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestStartUDBInstance ------------------
func TestStartUDBInstance(t *testing.T) {
	r := &StartUDBInstance{Region: "cn-east-01",
		DBId: "00f9868c-c7f5-4852-9eac-d200b678f0e1",
	}
	cmp := `https://api.ucloud.cn/?Action=StartUDBInstance&Region=cn-east-01&DBId=00f9868c-c7f5-4852-9eac-d200b678f0e1`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestStopUDBInstance ------------------
func TestStopUDBInstance(t *testing.T) {
	r := &StopUDBInstance{Region: "cn-east-01",
		DBId: "00f9868c-c7f5-4852-9eac-d200b678f0e1",
	}
	cmp := `https://api.ucloud.cn/?Action=StopUDBInstance&Region=cn-east-01&DBId=00f9868c-c7f5-4852-9eac-d200b678f0e1`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestUpdateUDBParamGroup ------------------
func TestUpdateUDBParamGroup(t *testing.T) {
	r := &UpdateUDBParamGroup{
		Region:  "cn-north-01",
		Value:   "2000",
		GroupId: 2,
		Key:     "back_log",
	}
	cmp := `https://api.ucloud.cn/?Region=cn-north-01&Action=UpdateUDBParamGroup&GroupId=2&Key=back_log&Value=2000`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}

// ---------------- TestUploadUDBParamGroup ------------------
func TestUploadUDBParamGroup(t *testing.T) {
	r := &UploadUDBParamGroup{Description: "xxx-config",
		GroupName: "mysql-xxx-config",
		Region:    "cn-east-01",
		Content:   `W2NsaWVudF0NCiBwb3J0ID0gMzMwNg0KIHNvY2tldCA9IC90bXAvbXlzcWwuc29jaw0KIFtteWlzYW1jaGtdDQoga2V5X2J1ZmZlcl9zaXplID0gMjBNDQogcmVhZF9idWZmZXIgPSAyTQ0KIHNvcnRfYnVmZmVyX3NpemUgPSAyME0NCiB3cml0ZV9idWZmZXIgPSAyTQ0KIFtteXNxbF0NCiBuby1hdXRvLXJlaGFzaA0KIFtteXNxbGRdDQogYmFja19sb2cgPSAyMDAwDQogYmFzZWRpciA9IC9vcHQvdWRiL3Byb2dyYW0vbXlzcWwvbXlzcWwtNS42LjIwDQogYmluZC1hZGRyZXNzID0gMTAuNC40LjE2OA0KIGJpbmxvZy1mb3JtYXQgPSBNSVhFRA0KIGNoYXJhY3Rlcl9zZXRfc2VydmVyID0gdXRmOA0KIGRhdGFkaXIgPSAvb3B0L3VkYi9pbnN0YW5jZS9teXNxbC01LjYvMTdlMzRmNjktODRkZi00ZGZlLTkzMTgtNGY1YTVkODI0NDcyL2RhdGENCiBldmVudF9zY2hlZHVsZXIgPSBPTg0KIGV4cGlyZV9sb2dzX2RheXMgPSA3DQogZ2VuZXJhbC1sb2ctZmlsZSA9IC9vcHQvdWRiL2luc3RhbmNlL215c3FsLTUuNi8xN2UzNGY2OS04NGRmLTRkZmUtOTMxOC00ZjVhNWQ4MjQ0NzIvbG9nL215c3FsZC5sb2cNCiBpbm5vZGJfYnVmZmVyX3Bvb2xfc2l6ZSA9IDYyOTE0NTYwMA0KIGlubm9kYl9kYXRhX2ZpbGVfcGF0aCA9IGliZGF0YTE6MTAwTTphdXRvZXh0ZW5kDQogaW5ub2RiX2RhdGFfaG9tZV9kaXIgPSAvb3B0L3VkYi9pbnN0YW5jZS9teXNxbC01LjYvMTdlMzRmNjktODRkZi00ZGZlLTkzMTgtNGY1YTVkODI0NDcyL2RhdGENCiBpbm5vZGJfZmlsZV9wZXJfdGFibGUgPSAxDQogaW5ub2RiX2ZsdXNoX2xvZ19hdF90cnhfY29tbWl0ID0gMg0KIGlubm9kYl9mbHVzaF9tZXRob2QgPSBPX0RJUkVDVA0KIGlubm9kYl9pb19jYXBhY2l0eSA9IDIwMDANCiBpbm5vZGJfbG9nX2J1ZmZlcl9zaXplID0gODM4ODYwOA0KIGlubm9kYl9sb2dfZmlsZV9zaXplID0gNTEyTQ0KIGlubm9kYl9sb2dfZmlsZXNfaW5fZ3JvdXAgPSAyDQogaW5ub2RiX2xvZ19ncm91cF9ob21lX2RpciA9IC9vcHQvdWRiL2luc3RhbmNlL215c3FsLTUuNi8xN2UzNGY2OS04NGRmLTRkZmUtOTMxOC00ZjVhNWQ4MjQ0NzIvZGF0YQ0KIGlubm9kYl9tYXhfZGlydHlfcGFnZXNfcGN0ID0gNTANCiBpbm5vZGJfb3Blbl9maWxlcyA9IDEwMjQNCiBpbm5vZGJfcmVhZF9pb190aHJlYWRzID0gOA0KIGlubm9kYl90aHJlYWRfY29uY3VycmVuY3kgPSAyMA0KIGlubm9kYl93cml0ZV9pb190aHJlYWRzID0gOA0KIGtleV9idWZmZXJfc2l6ZSA9IDMzNTU0NDMyDQogbG9jYWxfaW5maWxlID0gMA0KIGxvZy1iaW4gPSAvb3B0L3VkYi9pbnN0YW5jZS9teXNxbC01LjYvMTdlMzRmNjktODRkZi00ZGZlLTkzMTgtNGY1YTVkODI0NDcyL2JpbmxvZy9teXNxbC1iaW4ubG9nDQogbG9nLWVycm9yID0gL29wdC91ZGIvaW5zdGFuY2UvbXlzcWwtNS42LzE3ZTM0ZjY5LTg0ZGYtNGRmZS05MzE4LTRmNWE1ZDgyNDQ3Mi9sb2cvbXlzcWxkLmxvZw0KbG9nX2Jpbl90cnVzdF9mdW5jdGlvbl9jcmVhdG9ycyA9IDENCiBsb2dfb3V0cHV0ID0gVEFCTEUNCiBsb25nX3F1ZXJ5X3RpbWUgPSAzDQogbWF4X2FsbG93ZWRfcGFja2V0ID0gMTY3NzcyMTYNCiBtYXhfY29ubmVjdF9lcnJvcnMgPSAxMDAwMDAwDQogbWF4X2Nvbm5lY3Rpb25zID0gMjAwMA0KIG15aXNhbV9zb3J0X2J1ZmZlcl9zaXplID0gODM4ODYwOA0KIG5ldF9idWZmZXJfbGVuZ3RoID0gODE5Mg0KIHBpZC1maWxlID0gL29wdC91ZGIvaW5zdGFuY2UvbXlzcWwtNS42LzE3ZTM0ZjY5LTg0ZGYtNGRmZS05MzE4LTRmNWE1ZDgyNDQ3Mi9teXNxbGQucGlkDQogcG9ydCA9IDMzMDYNCiBxdWVyeV9jYWNoZV9zaXplID0gMTY3NzcyMTYNCiByZWFkX2J1ZmZlcl9zaXplID0gMjYyMTQ0DQogcmVhZF9ybmRfYnVmZmVyX3NpemUgPSA1MjQyODgNCiByZWxheS1sb2cgPSAvb3B0L3VkYi9pbnN0YW5jZS9teXNxbC01LjYvMTdlMzRmNjktODRkZi00ZGZlLTkzMTgtNGY1YTVkODI0NDcyL3JlbGF5bG9nL215c3FsLXJlbGF5LmxvZw0KIHNlY3VyZS1maWxlLXByaXYgPSAvb3B0L3VkYi9pbnN0YW5jZS9teXNxbC01LjYvMTdlMzRmNjktODRkZi00ZGZlLTkzMTgtNGY1YTVkODI0NDcyL3RtcA0KIHNlcnZlci1pZCA9IDE2ODAzNTQ5Ng0KIHNraXAtc2xhdmUtc3RhcnQNCiBza2lwX25hbWVfcmVzb2x2ZQ0KIHNsYXZlLWxvYWQtdG1wZGlyID0gL29wdC91ZGIvaW5zdGFuY2UvbXlzcWwtNS42LzE3ZTM0ZjY5LTg0ZGYtNGRmZS05MzE4LTRmNWE1ZDgyNDQ3Mi90bXANCiBzbG93LXF1ZXJ5LWxvZy1maWxlID0gL29wdC91ZGIvaW5zdGFuY2UvbXlzcWwtNS42LzE3ZTM0ZjY5LTg0ZGYtNGRmZS05MzE4LTRmNWE1ZDgyNDQ3Mi9sb2cvbXlzcWwtc2xvdy5sb2cNCiBzbG93X3F1ZXJ5X2xvZyA9IDENCiBzb2NrZXQgPSAvb3B0L3VkYi9pbnN0YW5jZS9teXNxbC01LjYvMTdlMzRmNjktODRkZi00ZGZlLTkzMTgtNGY1YTVkODI0NDcyL215c3FsZC5zb2NrDQogc29ydF9idWZmZXJfc2l6ZSA9IDUyNDI4OA0KIHN5bmNfYmlubG9nID0gMQ0KIHRhYmxlX29wZW5fY2FjaGUgPSAxMjgNCiB0aHJlYWRfY2FjaGVfc2l6ZSA9IDUwDQogdG1wZGlyID0gL29wdC91ZGIvaW5zdGFuY2UvbXlzcWwtNS42LzE3ZTM0ZjY5LTg0ZGYtNGRmZS05MzE4LTRmNWE1ZDgyNDQ3Mi90bXANCiB1c2VyID0gbXlzcWwNCiBbbXlzcWxkX3NhZmVdDQogbG9nLWVycm9yID0gL29wdC91ZGIvaW5zdGFuY2UvbXlzcWwtNS42LzE3ZTM0ZjY5LTg0ZGYtNGRmZS05MzE4LTRmNWE1ZDgyNDQ3Mi9sb2cvbXlzcWxkLmxvZw0KIHBpZC1maWxlID0gL29wdC91ZGIvaW5zdGFuY2UvbXlzcWwtNS42LzE3ZTM0ZjY5LTg0ZGYtNGRmZS05MzE4LTRmNWE1ZDgyNDQ3Mi9teXNxbGQucGlkDQogW215c3FsZHVtcF0NCiBtYXhfYWxsb3dlZF9wYWNrZXQgPSAxNk0NCiBxdWljaw0KIFtteXNxbGhvdGNvcHldDQogaW50ZXJhY3RpdmUtdGltZW91dA0KIC9vcHQvdWRiL2luc3RhbmNlL215c3E=`,
		DBTypeId:  "mysql-5.1",
	}
	cmp := `https://api.ucloud.cn/?Action=UploadUDBParamGroup&DBTypeId=mysql-5.1&GroupName=mysql-xxx-config&Description=xxx-config&Content=W2NsaWVudF0NCiBwb3J0ID0gMzMwNg0KIHNvY2tldCA9IC90bXAvbXlzcWwuc29jaw0KIFtteWlzYW1jaGtdDQoga2V5X2J1ZmZlcl9zaXplID0gMjBNDQogcmVhZF9idWZmZXIgPSAyTQ0KIHNvcnRfYnVmZmVyX3NpemUgPSAyME0NCiB3cml0ZV9idWZmZXIgPSAyTQ0KIFtteXNxbF0NCiBuby1hdXRvLXJlaGFzaA0KIFtteXNxbGRdDQogYmFja19sb2cgPSAyMDAwDQogYmFzZWRpciA9IC9vcHQvdWRiL3Byb2dyYW0vbXlzcWwvbXlzcWwtNS42LjIwDQogYmluZC1hZGRyZXNzID0gMTAuNC40LjE2OA0KIGJpbmxvZy1mb3JtYXQgPSBNSVhFRA0KIGNoYXJhY3Rlcl9zZXRfc2VydmVyID0gdXRmOA0KIGRhdGFkaXIgPSAvb3B0L3VkYi9pbnN0YW5jZS9teXNxbC01LjYvMTdlMzRmNjktODRkZi00ZGZlLTkzMTgtNGY1YTVkODI0NDcyL2RhdGENCiBldmVudF9zY2hlZHVsZXIgPSBPTg0KIGV4cGlyZV9sb2dzX2RheXMgPSA3DQogZ2VuZXJhbC1sb2ctZmlsZSA9IC9vcHQvdWRiL2luc3RhbmNlL215c3FsLTUuNi8xN2UzNGY2OS04NGRmLTRkZmUtOTMxOC00ZjVhNWQ4MjQ0NzIvbG9nL215c3FsZC5sb2cNCiBpbm5vZGJfYnVmZmVyX3Bvb2xfc2l6ZSA9IDYyOTE0NTYwMA0KIGlubm9kYl9kYXRhX2ZpbGVfcGF0aCA9IGliZGF0YTE6MTAwTTphdXRvZXh0ZW5kDQogaW5ub2RiX2RhdGFfaG9tZV9kaXIgPSAvb3B0L3VkYi9pbnN0YW5jZS9teXNxbC01LjYvMTdlMzRmNjktODRkZi00ZGZlLTkzMTgtNGY1YTVkODI0NDcyL2RhdGENCiBpbm5vZGJfZmlsZV9wZXJfdGFibGUgPSAxDQogaW5ub2RiX2ZsdXNoX2xvZ19hdF90cnhfY29tbWl0ID0gMg0KIGlubm9kYl9mbHVzaF9tZXRob2QgPSBPX0RJUkVDVA0KIGlubm9kYl9pb19jYXBhY2l0eSA9IDIwMDANCiBpbm5vZGJfbG9nX2J1ZmZlcl9zaXplID0gODM4ODYwOA0KIGlubm9kYl9sb2dfZmlsZV9zaXplID0gNTEyTQ0KIGlubm9kYl9sb2dfZmlsZXNfaW5fZ3JvdXAgPSAyDQogaW5ub2RiX2xvZ19ncm91cF9ob21lX2RpciA9IC9vcHQvdWRiL2luc3RhbmNlL215c3FsLTUuNi8xN2UzNGY2OS04NGRmLTRkZmUtOTMxOC00ZjVhNWQ4MjQ0NzIvZGF0YQ0KIGlubm9kYl9tYXhfZGlydHlfcGFnZXNfcGN0ID0gNTANCiBpbm5vZGJfb3Blbl9maWxlcyA9IDEwMjQNCiBpbm5vZGJfcmVhZF9pb190aHJlYWRzID0gOA0KIGlubm9kYl90aHJlYWRfY29uY3VycmVuY3kgPSAyMA0KIGlubm9kYl93cml0ZV9pb190aHJlYWRzID0gOA0KIGtleV9idWZmZXJfc2l6ZSA9IDMzNTU0NDMyDQogbG9jYWxfaW5maWxlID0gMA0KIGxvZy1iaW4gPSAvb3B0L3VkYi9pbnN0YW5jZS9teXNxbC01LjYvMTdlMzRmNjktODRkZi00ZGZlLTkzMTgtNGY1YTVkODI0NDcyL2JpbmxvZy9teXNxbC1iaW4ubG9nDQogbG9nLWVycm9yID0gL29wdC91ZGIvaW5zdGFuY2UvbXlzcWwtNS42LzE3ZTM0ZjY5LTg0ZGYtNGRmZS05MzE4LTRmNWE1ZDgyNDQ3Mi9sb2cvbXlzcWxkLmxvZw0KbG9nX2Jpbl90cnVzdF9mdW5jdGlvbl9jcmVhdG9ycyA9IDENCiBsb2dfb3V0cHV0ID0gVEFCTEUNCiBsb25nX3F1ZXJ5X3RpbWUgPSAzDQogbWF4X2FsbG93ZWRfcGFja2V0ID0gMTY3NzcyMTYNCiBtYXhfY29ubmVjdF9lcnJvcnMgPSAxMDAwMDAwDQogbWF4X2Nvbm5lY3Rpb25zID0gMjAwMA0KIG15aXNhbV9zb3J0X2J1ZmZlcl9zaXplID0gODM4ODYwOA0KIG5ldF9idWZmZXJfbGVuZ3RoID0gODE5Mg0KIHBpZC1maWxlID0gL29wdC91ZGIvaW5zdGFuY2UvbXlzcWwtNS42LzE3ZTM0ZjY5LTg0ZGYtNGRmZS05MzE4LTRmNWE1ZDgyNDQ3Mi9teXNxbGQucGlkDQogcG9ydCA9IDMzMDYNCiBxdWVyeV9jYWNoZV9zaXplID0gMTY3NzcyMTYNCiByZWFkX2J1ZmZlcl9zaXplID0gMjYyMTQ0DQogcmVhZF9ybmRfYnVmZmVyX3NpemUgPSA1MjQyODgNCiByZWxheS1sb2cgPSAvb3B0L3VkYi9pbnN0YW5jZS9teXNxbC01LjYvMTdlMzRmNjktODRkZi00ZGZlLTkzMTgtNGY1YTVkODI0NDcyL3JlbGF5bG9nL215c3FsLXJlbGF5LmxvZw0KIHNlY3VyZS1maWxlLXByaXYgPSAvb3B0L3VkYi9pbnN0YW5jZS9teXNxbC01LjYvMTdlMzRmNjktODRkZi00ZGZlLTkzMTgtNGY1YTVkODI0NDcyL3RtcA0KIHNlcnZlci1pZCA9IDE2ODAzNTQ5Ng0KIHNraXAtc2xhdmUtc3RhcnQNCiBza2lwX25hbWVfcmVzb2x2ZQ0KIHNsYXZlLWxvYWQtdG1wZGlyID0gL29wdC91ZGIvaW5zdGFuY2UvbXlzcWwtNS42LzE3ZTM0ZjY5LTg0ZGYtNGRmZS05MzE4LTRmNWE1ZDgyNDQ3Mi90bXANCiBzbG93LXF1ZXJ5LWxvZy1maWxlID0gL29wdC91ZGIvaW5zdGFuY2UvbXlzcWwtNS42LzE3ZTM0ZjY5LTg0ZGYtNGRmZS05MzE4LTRmNWE1ZDgyNDQ3Mi9sb2cvbXlzcWwtc2xvdy5sb2cNCiBzbG93X3F1ZXJ5X2xvZyA9IDENCiBzb2NrZXQgPSAvb3B0L3VkYi9pbnN0YW5jZS9teXNxbC01LjYvMTdlMzRmNjktODRkZi00ZGZlLTkzMTgtNGY1YTVkODI0NDcyL215c3FsZC5zb2NrDQogc29ydF9idWZmZXJfc2l6ZSA9IDUyNDI4OA0KIHN5bmNfYmlubG9nID0gMQ0KIHRhYmxlX29wZW5fY2FjaGUgPSAxMjgNCiB0aHJlYWRfY2FjaGVfc2l6ZSA9IDUwDQogdG1wZGlyID0gL29wdC91ZGIvaW5zdGFuY2UvbXlzcWwtNS42LzE3ZTM0ZjY5LTg0ZGYtNGRmZS05MzE4LTRmNWE1ZDgyNDQ3Mi90bXANCiB1c2VyID0gbXlzcWwNCiBbbXlzcWxkX3NhZmVdDQogbG9nLWVycm9yID0gL29wdC91ZGIvaW5zdGFuY2UvbXlzcWwtNS42LzE3ZTM0ZjY5LTg0ZGYtNGRmZS05MzE4LTRmNWE1ZDgyNDQ3Mi9sb2cvbXlzcWxkLmxvZw0KIHBpZC1maWxlID0gL29wdC91ZGIvaW5zdGFuY2UvbXlzcWwtNS42LzE3ZTM0ZjY5LTg0ZGYtNGRmZS05MzE4LTRmNWE1ZDgyNDQ3Mi9teXNxbGQucGlkDQogW215c3FsZHVtcF0NCiBtYXhfYWxsb3dlZF9wYWNrZXQgPSAxNk0NCiBxdWljaw0KIFtteXNxbGhvdGNvcHldDQogaW50ZXJhY3RpdmUtdGltZW91dA0KIC9vcHQvdWRiL2luc3RhbmNlL215c3E%3D&DBTypeId=mysql-5.5&Description=test&GroupName=test1222&Region=cn-east-01`

	if err := FakeGetAndCmp(r, cmp); err != nil {
		t.Fatal(err)
	}
}
