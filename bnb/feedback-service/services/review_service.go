package services

import (
	"errors"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/dtos"
	grpcclient "github.com/rruzicic/globetrotter/bnb/feedback-service/grpc_client"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/models"
	"github.com/rruzicic/globetrotter/bnb/feedback-service/repos"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateHostReview(hostReviewDTO dtos.CreateHostReviewDTO) (*models.HostReview, error) {
	userId, err := primitive.ObjectIDFromHex(hostReviewDTO.UserId)
	if err != nil {
		return nil, err
	}

	hostId, err := primitive.ObjectIDFromHex(hostReviewDTO.HostId)
	if err != nil {
		return nil, err
	}

	hostReview := models.HostReview{
		UserId: &userId,
		HostId: &hostId,
		Rating: hostReviewDTO.Rating,
	}

	//TODO check if user had a previous reservation with this host.
	finishedReservations, err := grpcclient.GetFinishedReservationsByUser(userId.Hex())
	if err != nil {
		return nil, err
	}

	accommodationIds := []string{}
	for _, reservation := range finishedReservations {
		accommodationIds = append(accommodationIds, reservation.AccommodationId)
	}

	pastHostsIds, err := grpcclient.GetPastHostsByAccommodations(accommodationIds)
	if err != nil {
		return nil, err
	}
	hasUserBeenToHosts := false
	for _, pastHost := range pastHostsIds {
		if pastHost.HostId == hostId.Hex() {
			hasUserBeenToHosts = true
		}
	}

	//Publish an event to the account service
	conn := Conn()
	defer conn.Close()

	err = conn.Publish("account-service", []byte("calculate avg rating"))
	if err != nil {
		log.Panic(err)
	}

	if hasUserBeenToHosts {
		return repos.CreateHostReview(hostReview)
	} else {
		return nil, errors.New("user has not been to any of the hosts accommodations before and therefore can't review him")
	}
}

func GetHostReviewById(id string) (*models.HostReview, error) {
	return repos.GetHostReviewById(id)
}

func GetHostReviewsByUserId(id string) ([]models.HostReview, error) {
	return repos.GetHostReviewsByUserId(id)
}

func GetHostReviewsByHostId(id string) ([]models.HostReview, error) {
	return repos.GetHostReviewsByHostId(id)
}

func DeleteHostReview(id string) error {
	return repos.DeleteHostReview(id)
}

func UpdateHostReview(hostReviewDTO dtos.CreateHostReviewDTO) error {
	userId, err := primitive.ObjectIDFromHex(hostReviewDTO.UserId)
	if err != nil {
		return err
	}

	hostId, err := primitive.ObjectIDFromHex(hostReviewDTO.HostId)
	if err != nil {
		return err
	}

	hostReview := models.HostReview{
		UserId: &userId,
		HostId: &hostId,
		Rating: hostReviewDTO.Rating,
	}

	return repos.UpdateHostReview(hostReview)
}

//===================================================================================
//===================================================================================

func CreateAccommodationReview(accommodationReviewDTO dtos.CreateAccommodationReviewDTO) (*models.AccommodationReview, error) {
	userId, err := primitive.ObjectIDFromHex(accommodationReviewDTO.UserId)
	if err != nil {
		return nil, err
	}

	accommodationId, err := primitive.ObjectIDFromHex(accommodationReviewDTO.AccommodationId)
	if err != nil {
		return nil, err
	}

	accommodationReview := models.AccommodationReview{
		UserId:          &userId,
		AccommodationId: &accommodationId,
		Rating:          accommodationReviewDTO.Rating,
	}

	finishedReservations, err := grpcclient.GetFinishedReservationsByUser(userId.Hex())
	if err != nil {
		return nil, err
	}
	hasUserBeenToAccommodation := false
	for _, reservation := range finishedReservations {
		if accommodationId.Hex() == reservation.AccommodationId {
			hasUserBeenToAccommodation = true
			break
		}
	}

	if hasUserBeenToAccommodation {
		return repos.CreateAccommodationReview(accommodationReview)
	} else {
		return nil, errors.New("user has not been to accommodation before and therefore can't review it")
	}

}

func GetAccommodationtReviewById(id string) (*models.AccommodationReview, error) {
	return repos.GetAccommodationReviewById(id)
}

func GetAccommodationReviewsByUserId(id string) ([]models.AccommodationReview, error) {
	return repos.GetAccommodationReviewsByUserId(id)
}

func GetAccommodationReviewsByAccommodationId(id string) ([]models.AccommodationReview, error) {
	return repos.GetAccommodationReviewsByAccommodationId(id)
}

func DeleteAccommodationReview(id string) error {
	return repos.DeleteAccommodationReview(id)
}

func UpdateAccommodationReview(accommodationReviewDTO dtos.CreateAccommodationReviewDTO) error {
	userId, err := primitive.ObjectIDFromHex(accommodationReviewDTO.UserId)
	if err != nil {
		return err
	}

	accommodationId, err := primitive.ObjectIDFromHex(accommodationReviewDTO.AccommodationId)
	if err != nil {
		return err
	}

	accommodationReview := models.AccommodationReview{
		UserId:          &userId,
		AccommodationId: &accommodationId,
		Rating:          accommodationReviewDTO.Rating,
	}

	return repos.UpdateAccommodationReview(accommodationReview)
}

func Conn() *nats.Conn {
	conn, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
