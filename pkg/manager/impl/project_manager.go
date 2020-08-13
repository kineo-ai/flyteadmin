package impl

import (
	"context"

	"github.com/lyft/flyteadmin/pkg/common"
	"github.com/lyft/flyteadmin/pkg/manager/impl/validation"
	"github.com/lyft/flyteadmin/pkg/manager/interfaces"
	"github.com/lyft/flyteadmin/pkg/repositories"
	"github.com/lyft/flyteadmin/pkg/repositories/transformers"
	runtimeInterfaces "github.com/lyft/flyteadmin/pkg/runtime/interfaces"
	"github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"
)

type ProjectManager struct {
	db     repositories.RepositoryInterface
	config runtimeInterfaces.Configuration
}

var alphabeticalSortParam, _ = common.NewSortParameter(admin.Sort{
	Direction: admin.Sort_ASCENDING,
	Key:       "identifier",
})

func (m *ProjectManager) CreateProject(ctx context.Context, request admin.ProjectRegisterRequest) (
	*admin.ProjectRegisterResponse, error) {
	if err := validation.ValidateProjectRegisterRequest(request); err != nil {
		return nil, err
	}
	projectModel := transformers.CreateProjectModel(request.Project)
	err := m.db.ProjectRepo().Create(ctx, projectModel)
	if err != nil {
		return nil, err
	}

	return &admin.ProjectRegisterResponse{}, nil
}

func (m *ProjectManager) getDomains() []*admin.Domain {
	configDomains := m.config.ApplicationConfiguration().GetDomainsConfig()
	var domains = make([]*admin.Domain, len(*configDomains))
	for index, configDomain := range *configDomains {
		domains[index] = &admin.Domain{
			Id:   configDomain.ID,
			Name: configDomain.Name,
		}
	}
	return domains
}

func (m *ProjectManager) ListProjects(ctx context.Context, request admin.ProjectListRequest) (*admin.Projects, error) {
	projectModels, err := m.db.ProjectRepo().ListAll(ctx, alphabeticalSortParam)
	if err != nil {
		return nil, err
	}

	projects := transformers.FromProjectModels(projectModels, m.getDomains())
	return &admin.Projects{
		Projects: projects,
	}, nil
}

func (m *ProjectManager) UpdateProject(ctx context.Context, project admin.Project) (*admin.ProjectUpdateResponse, error) {
	var response *admin.ProjectUpdateResponse

	projectModel := transformers.CreateProjectModel(&project);
	_, err := m.db.ProjectRepo().UpdateProject(ctx, projectModel)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewProjectManager(db repositories.RepositoryInterface, config runtimeInterfaces.Configuration) interfaces.ProjectInterface {
	return &ProjectManager{
		db:     db,
		config: config,
	}
}
