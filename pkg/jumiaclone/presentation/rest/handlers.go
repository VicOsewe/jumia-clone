package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/domain/dao"
	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/domain/dto"
	"github.com/VicOsewe/Jumia-clone/pkg/jumiaclone/usecases/onboarding"
)

// Presentation represents the presentation layer contract
type Presentation interface {
	CreateUser() http.HandlerFunc
	LoginInByEMail() http.HandlerFunc
	CheckIfPhoneNumberExists() http.HandlerFunc
	VerifyPhoneNumber() http.HandlerFunc
	CheckIfEmailExists() http.HandlerFunc
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
		input := &dao.UserProfile{}
		DecodeJSONToTargetStruct(w, r, input)
		validatedInput, err := validateCreateUserInput(input)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		response, err := rs.onboardingUsecases.CreateUser(validatedInput)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		marshalled, err := json.Marshal(dto.APIResponseMessage{
			Message:    "user has been created successfully",
			StatusCode: http.StatusOK,
			Body:       response,
		},
		)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		RespondWithJSON(w, http.StatusOK, marshalled)
	}
}

func (rs *RestFulAPIs) CheckIfPhoneNumberExists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := &dao.UserProfile{}
		DecodeJSONToTargetStruct(w, r, input)
		if input.PhoneNumber == "" {
			RespondWithError(w, http.StatusBadRequest, fmt.Errorf("expected a phone number to be given but it was not supplied"))
			return
		}

		response, err := rs.onboardingUsecases.CheckIfPhoneNumberExists(input.PhoneNumber)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		marshalled, err := json.Marshal(response)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		RespondWithJSON(w, http.StatusOK, marshalled)
	}
}

func (rs *RestFulAPIs) CheckIfEmailExists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := &dao.UserProfile{}
		DecodeJSONToTargetStruct(w, r, input)
		if input.Email == "" {
			RespondWithError(w, http.StatusBadRequest, fmt.Errorf("expected a email to be given but it was not supplied"))
			return
		}

		response, err := rs.onboardingUsecases.CheckIfEmailExists(input.PhoneNumber)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err)
			return
		}
		marshalled, err := json.Marshal(response)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		RespondWithJSON(w, http.StatusOK, marshalled)
	}
}

func (rs *RestFulAPIs) VerifyPhoneNumber() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := &dao.OTPPayload{}
		DecodeJSONToTargetStruct(w, r, input)

		otp, err := validateOtpPayload(input)
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err)
			return
		}

		response, err := rs.onboardingUsecases.VerifyPhoneNumber(otp)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err)
			return
		}

		marshalled, err := json.Marshal(response)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		RespondWithJSON(w, http.StatusOK, marshalled)
	}
}

func (rs *RestFulAPIs) LoginInByEMail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := &dao.UserProfile{}
		DecodeJSONToTargetStruct(w, r, input)
		if input.Email == "" {
			RespondWithError(w, http.StatusBadRequest, fmt.Errorf("expected a email to be given but it was not supplied"))
			return
		}

		if input.PassWord == "" {
			RespondWithError(w, http.StatusBadRequest, fmt.Errorf("expected password to be given but it was not supplied"))
			return
		}

		response, err := rs.onboardingUsecases.SignInByEmail(input.Email, input.PassWord)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		marshalled, err := json.Marshal(response)
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err)
			return
		}
		RespondWithJSON(w, http.StatusOK, marshalled)
	}
}
