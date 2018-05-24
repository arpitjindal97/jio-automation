package main

import (
	"github.com/therecipe/qt/core"
)

func init() { CustomTableModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "CustomTableModel") }

const (
	IMEI = int(core.Qt__UserRole) + 1<<iota
	Distributor
	Retailer
	Head1
	Date
)

type TableItem struct {
	IMEI        string
	Distributor string
	Retailer    string
	Head1       string
	Date        string
}

type CustomTableModel struct {
	core.QAbstractTableModel

	_ func() `constructor:"init"`

	_ func()                          `signal:"remove,auto"`
	_ func(item TableItem)            `signal:"add,auto"`
	_ func(index int, item TableItem) `signal:"edit,auto"`
	_ func(filePath string)           `signal:"openfile,auto"`
	_ func()                          `signal:"start,auto"`

	modelData []TableItem
}

func (m *CustomTableModel) init() {
	m.ConnectRoleNames(m.roleNames)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectData(m.data)
}

func (m *CustomTableModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		IMEI:        core.NewQByteArray2("IMEI", -1),
		Distributor: core.NewQByteArray2("Distributor", -1),
		Retailer:    core.NewQByteArray2("Retailer", -1),
		Head1:       core.NewQByteArray2("Head1", -1),
		Date:        core.NewQByteArray2("Date", -1),
	}
}

func (m *CustomTableModel) rowCount(*core.QModelIndex) int {
	return len(m.modelData)
}

func (m *CustomTableModel) columnCount(*core.QModelIndex) int {
	return 5
}

func (m *CustomTableModel) data(index *core.QModelIndex, role int) *core.QVariant {
	item := m.modelData[index.Row()]
	switch role {
	case IMEI:
		return core.NewQVariant14(item.IMEI)
	case Distributor:
		return core.NewQVariant14(item.Distributor)
	case Retailer:
		return core.NewQVariant14(item.Retailer)
	case Head1:
		return core.NewQVariant14(item.Head1)
	case Date:
		return core.NewQVariant14(item.Date)

	}
	return core.NewQVariant()
}

func (m *CustomTableModel) remove() {
	if len(m.modelData) == 0 {
		return
	}
	m.BeginRemoveRows(core.NewQModelIndex(), len(m.modelData)-1, len(m.modelData)-1)
	m.modelData = m.modelData[:len(m.modelData)-1]
	m.EndRemoveRows()
}

func (m *CustomTableModel) add(item TableItem) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	m.modelData = append(m.modelData, item)
	m.EndInsertRows()
}

func (m *CustomTableModel) edit(index int, item TableItem) {
	if len(m.modelData) == 0 {
		return
	}
	m.modelData[index] = item
	m.DataChanged(m.Index(index, 0, core.NewQModelIndex()), m.Index(index, 1, core.NewQModelIndex()),
		[]int{IMEI, Distributor, Retailer, Head1, Date})
}

func (m *CustomTableModel) openfile(filePath string) {
	SetModalInstance(m)
	OpenFile(filePath)
}

func (m *CustomTableModel) start() {
	go StartProcess()
}
