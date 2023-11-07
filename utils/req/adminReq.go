package req

import (
	"peekabook/model/domain"
	"peekabook/model/schema"
	"peekabook/model/web"
)

func AdminCreateRequestToAdminDomain(request web.AdminCreateRequest) *domain.Admin {
	return &domain.Admin{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}

func AdminLoginRequestToAdminDomain(request web.AdminLoginRequest) *domain.Admin {
	return &domain.Admin{
		Email:    request.Email,
		Password: request.Password,
	}
}

func AdminUpdateRequestToAdminDomain(request web.AdminUpdateRequest) *domain.Admin {
	return &domain.Admin{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}

func AdminResetPasswordRequestToAdminDomain(request web.AdminResetPasswordRequest) *domain.Admin {
	return &domain.Admin{
		Password: request.NewPassword,
	}
}

func AdminDomaintoAdminSchema(request domain.Admin) *schema.Admin {
	return &schema.Admin{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
}
