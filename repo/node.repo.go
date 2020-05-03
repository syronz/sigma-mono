package repo

import (
	"sigmamono/internal/core"
	"sigmamono/internal/param"
	"sigmamono/internal/search"
	"sigmamono/internal/types"
	"sigmamono/model"

	"github.com/jinzhu/gorm"
)

// NodeRepo for injecting engine
type NodeRepo struct {
	Engine *core.Engine
}

// ProvideNodeRepo is used in wire
func ProvideNodeRepo(engine *core.Engine) NodeRepo {
	return NodeRepo{Engine: engine}
}

// FindByID for node
func (p *NodeRepo) FindByID(id types.RowID) (node model.Node, err error) {
	err = p.Engine.DB.First(&node, id.ToUint64()).Error
	return
}

// List of nodes
func (p *NodeRepo) List(params param.Param) (nodes []model.Node, err error) {
	columns, err := model.Node{}.Columns(params.Select)
	if err != nil {
		return
	}

	err = p.Engine.DB.Select(columns).
		Joins("INNER JOIN companies on companies.id = nodes.company_id").
		Where(search.Parse(params, model.Node{}.Pattern())).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&nodes).Error

	return
}

// Count of nodes
func (p *NodeRepo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table("nodes").
		Select(params.Select).
		Where("deleted_at is null").
		Where(search.Parse(params, model.Node{}.Pattern())).
		Count(&count).Error
	return
}

// Update NodeRepo
func (p *NodeRepo) Update(node model.Node) (u model.Node, err error) {
	err = p.Engine.DB.Save(&node).Error
	p.Engine.DB.Where("id = ?", node.ID).Find(&u)
	return
}

// Create NodeRepo
func (p *NodeRepo) Create(node model.Node) (u model.Node, err error) {
	err = p.Engine.DB.Create(&node).Scan(&u).Error
	return
}

// LastNode of node table
func (p *NodeRepo) LastNode(companyID types.RowID) (node model.Node, err error) {
	err = p.Engine.DB.Unscoped().Where("company_id LIKE ?", companyID).Last(&node).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}
	return
}

// Delete node
func (p *NodeRepo) Delete(node model.Node) (err error) {
	err = p.Engine.DB.Delete(&node).Error
	return
}

// HardDelete node
func (p *NodeRepo) HardDelete(node model.Node) (err error) {
	err = p.Engine.DB.Unscoped().Delete(&node).Error
	return
}
