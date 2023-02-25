package rest

import (
	"encoding/json"
	"net/http"

	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/domain/dao"
	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/usecases/onboarding"
)

// Presentation represents the presentation layer contract
type Presentation interface {
	CreateUser() http.HandlerFunc
}

// RestFulAPIs sets up RESTFUL APIs with all necessary dependencies
type RestFulAPIs struct {
	onboardingUsecases onboarding.OnboardingUsecase
}

func NewRestFulAPIs(on onboarding.OnboardingUsecase) *RestFulAPIs {
	rs := &RestFulAPIs{
		onboardingUsecases: on,
	}
	return rs
}

func (rs *RestFulAPIs) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := &dao.User{}
		DecodeJSONToTargetStruct(w, r, input)
		validatedInput, err := validateCreateUserInput(input)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		resp, err := rs.onboardingUsecases.CreateUser(validatedInput)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		marshalled, err := json.Marshal(resp)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		RespondWithJSON(w, http.StatusOK, marshalled)

	}
}
