package handler

func Register(r *Router) {
	r.GET("/skills/:id", r.SkillHandler.GetSkillByID)
	r.GET("/skills/", r.SkillHandler.GetSkillByKeyword)
	r.GET("/skills/bulk", r.SkillHandler.GetBulkSkill)
	r.PUT("/skills/:id", r.SkillHandler.UpdateSkillByID)
}
