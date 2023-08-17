package main

import (
	"fmt"
	"github.com/sheelendar/src/allen_batch_allocations/data_storage"
	"github.com/sheelendar/src/allen_batch_allocations/handler/handler"
	"github.com/sheelendar/src/allen_batch_allocations/utils"
)

func main() {
	handler.Register("Akhilesh", utils.MALE, utils.STUDENT)
	handler.Register("Komal", utils.FEMALE, utils.STUDENT)
	handler.Register("Rajnish", utils.MALE, utils.STUDENT)
	handler.Register("Mayuri", utils.FEMALE, utils.STUDENT)
	handler.Register("Ram", utils.MALE, utils.STUDENT)

	handler.Enroll("Akhilesh", "IIT")
	handler.Enroll("Komal", "IIT")
	handler.Enroll("Rajnish", "NEET")
	handler.Enroll("Mayuri", "IIT")
	handler.Enroll("Ram", "IIT")

	handler.Register("Kamesh", utils.MALE, utils.ADMIN)
	handler.Register("M", utils.MALE, utils.ADMIN)

	handler.CreateBatch("Kamesh", "IIT", utils.MORNING, 3, 1)
	handler.CreateBatch("Kamesh", "NEET", utils.EVENING, 2, 20)
	handler.CreateBatch("Kamesh", "IIT", utils.EVENING, 3, 21)

	handler.AllocateBatch("Kamesh", "Komal", utils.AllocationGenderPriority)
	handler.AllocateBatch("Kamesh", "Mayuri", utils.AllocationHigerCapa)
	handler.AllocateBatch("Kamesh", "Akhilesh", utils.AllocationHigerCapa)
	handler.AllocateBatch("Kamesh", "Rajnish", utils.AllocationHigerCapa)
	handler.AllocateBatch("Kamesh", "Ram", utils.AllocationHigerCapa)

	for name, student := range data_storage.StudentList {
		fmt.Println("Name: ", name)
		fmt.Println("stream: ", student.Stream, " Roll: ", student.Roll, " Gender: ", student.Gender, " batchID ", student.BatchID)
	}

	fmt.Println()
	fmt.Println()
	for name, admin := range data_storage.AdminList {
		fmt.Println("Name: ", name)
		fmt.Println(" Roll: ", admin.Roll, " Gender: ", admin.Gender)
	}

	fmt.Println()
	n := len(data_storage.Batches)

	for i := 0; i < n; i++ {
		fmt.Print(" Name ", data_storage.Batches[i].Name)
		fmt.Print(" BatchID ", data_storage.Batches[i].BatchID)
		fmt.Print(" Capacity ", data_storage.Batches[i].Capacity)
		fmt.Print(" CurrCapacity ", data_storage.Batches[i].CurrCapacity)
		fmt.Println()

	}
}
