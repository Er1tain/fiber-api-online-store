package admin

func (controller Admin) openAdminPanel() {
	controller.service.Static("/", "src/admin/static")
}
