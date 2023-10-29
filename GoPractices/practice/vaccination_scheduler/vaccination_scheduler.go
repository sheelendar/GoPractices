// Our goal is to implement and demonstrate a service for connecting vaccine patients with providers. Patients should be able to browse available appointments, make or cancel a reservation, and review the details of their reservation. Providers need to manage their available appointments and review their daily appointment reservations.

// Patients may only have one reservation in total at any given time (not one per provider).

// All vaccination appointments last 30 minutes and start on the hour or half hour. Providers may otherwise create appointments at any time of day.

// The desired API is provided, and may be adapted to your preferred programming language.

package main

import (
	"errors"
	"fmt"
	"time"
)

type VaccineScheduler struct {
	slot   map[string]map[string]map[time.Time]bool
	pspMap map[string]ProviderAppointment
}

type Provider struct {
	ProviderID string
}
type ProviderAppointment struct {
	PatientID       string
	AppointmentTime time.Time
	Provider        Provider
}

// ScheduleAppointment reserves a patient appointment with this provider and
// appointment time
//
// Returns:
//
//	An error if this appointment does not exist or this patient already has
//	an appointment
func (s *VaccineScheduler) ScheduleAppointment(patientID string, providerID string, appointmentTime time.Time) error {
	timeSlot, ok := s.slot[providerID]
	if !ok {
		return errors.New("No such provider")
	}
	date := appointmentTime.Format("2006-01-02")
	slots, ok := timeSlot[date]
	if !ok {
		return errors.New("There is no available slot for this time")
	}
	if !slots[appointmentTime] {
		return errors.New("This slot is already booked")
	}
	slots[appointmentTime] = false
	s.pspMap[patientID] = ProviderAppointment{PatientID: patientID, Provider: Provider{ProviderID: providerID}, AppointmentTime: appointmentTime}
	return nil
}

// CancelAppointment cancels an existing appointment for this patient. If this
// patient has no appointment, do nothing.
func (s *VaccineScheduler) CancelAppointment(patientID string) {
	pa, ok := s.pspMap[patientID]
	if !ok {
		fmt.Println("No appointment for this patient")
	}
	delete(s.pspMap, patientID)
	date := pa.AppointmentTime.Format("2006-01-02")
	s.slot[pa.Provider.ProviderID][date][pa.AppointmentTime] = true

}

// GetPatientAppointment gets this patient's appointment information
//
// Returns:
//
//	This patient's appointment time and provider ID, or
//	empty values if this patient has no scheduled appointment
func (s *VaccineScheduler) GetPatientAppointment(patientID string) (time.Time, string) {
	if v, ok := s.pspMap[patientID]; ok {
		return v.AppointmentTime, v.Provider.ProviderID
	}
	return time.Now(), ""
}

// GetAvailableAppointments gets open appointments on this day
// for patients to browse (day contains a calendar date)
//
// Returns:
//
//	A mapping of appointment time to list of provider IDs indicating which
//	providers have available appointments for each appointment time on this
//	day
func (s *VaccineScheduler) GetAvailableAppointments(day time.Time) map[time.Time][]string {
	result := make(map[time.Time][]string)
	dayDate := day.Format("2006-01-02")

	for provider, daySlots := range s.slot {
		slots := daySlots[dayDate]
		for timeStam, ok := range slots {
			if ok {
				result[timeStam] = append(result[timeStam], provider)
			}
		}
	}
	return result
}

// AddAppointment makes a new appointment with this provider available
//
// Returns:
//
//	An error if this provider already has an appointment at this time
func (s *VaccineScheduler) AddAppointment(providerID string, appointmentTime time.Time) error {
	date := appointmentTime.Format("2006-01-02")
	daySlotMap, ok := s.slot[providerID]
	if !ok {
		s.slot[providerID] = make(map[string]map[time.Time]bool)
		newMap := s.slot[providerID]
		newMap[date] = make(map[time.Time]bool)
		newMap[date][appointmentTime] = true
		return nil
	}
	if _, ok := daySlotMap[date][appointmentTime]; ok {
		return errors.New("slot is already added")
	}
	daySlotMap[date][appointmentTime] = true
	return nil
}

// RemoveAppointment removes an available appointment for a provider at this
// time. If this provider has no appointment at this time, do nothing.
func (s *VaccineScheduler) RemoveAppointment(providerID string, appointmentTime time.Time) {
	daySlot, ok := s.slot[providerID]

	if !ok {
		fmt.Println("No such provider")
		return
	}
	date := appointmentTime.Format("2006-01-02")
	slots, ok := daySlot[date]
	if !ok {
		fmt.Println("No such day exits for the slot")
	}
	_, is := slots[appointmentTime]
	if !is {
		fmt.Println("No such appointment")
	}
	delete(slots, appointmentTime)
}

// GetProviderSchedule gets this provider's patient schedule for this day
//
// Returns:
//
//	A list of ProviderAppointment structs (containing appointment times and
//	patient IDs), sorted by appointment time, which represents the patient
//	schedule for this provider on this day
func (s *VaccineScheduler) GetProviderSchedule(providerID string, day time.Time) []ProviderAppointment {
	return nil
}

func main() {
	vaccineScheduler := &VaccineScheduler{
		slot:   make(map[string]map[string]map[time.Time]bool),
		pspMap: make(map[string]ProviderAppointment),
	}
	stime := time.Now()

	vaccineScheduler.AddAppointment("10001", stime)
	fmt.Println("sheelendar")
	for providerID, dateSlotMap := range vaccineScheduler.slot {
		fmt.Println(providerID)
		fmt.Println(dateSlotMap)
	}
	err := vaccineScheduler.AddAppointment("10001", stime)
	fmt.Println(err)

	vaccineScheduler.AddAppointment("10002", stime)
	secondTime := stime.Add(time.Hour * 1)
	vaccineScheduler.AddAppointment("10001", secondTime)
	vaccineScheduler.AddAppointment("10002", secondTime)

	for providerID, dateSlotMap := range vaccineScheduler.slot {
		fmt.Println(providerID)
		fmt.Println(dateSlotMap)
	}

	vaccineScheduler.ScheduleAppointment("12", "10001", secondTime)
	res := vaccineScheduler.GetAvailableAppointments(secondTime)
	for time, providers := range res {
		fmt.Print(time, "  - ")
		fmt.Println(providers)
	}
	for providerID, dateSlotMap := range vaccineScheduler.slot {
		fmt.Println(providerID)
		fmt.Println(dateSlotMap)
	}
}
