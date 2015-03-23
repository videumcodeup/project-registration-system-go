package modelhelper

import (
	"github.com/videumcodeup/project-registration-system-go/infrastructure/entities"
	"github.com/videumcodeup/project-registration-system-go/models"
)

func MapToCodeUpViewModel(codeUpDbEntity entities.CodeUp) *models.CodeUpViewModel{
  var vm = new (models.CodeUpViewModel)
  vm.Id = codeUpDbEntity.Id
  vm.Title = codeUpDbEntity.Title
  vm.StartTimeText = codeUpDbEntity.StartTime.Format("Mon Jan _2 15:04:05 2006")
  vm.EndTimeText = codeUpDbEntity.EndTime.Format("Mon Jan _2 15:04:05 2006")
  return vm
}

//func MapToCodeUpViewModelList(codeUpDbEntityList *[]entities.CodeUp) []models.CodeUpViewModel{

  //viewModelList := make([]models.CodeUpViewModel, len(codeUpDbEntityList))
/*
  for i, v := range codeUpDbEntityList {
    viewModelList[i] = MapToCodeUpViewModel(entities.CodeUp(v))
  }
  */

//  return viewModelList
//}
