package convertor

//import (
//	"github.com/VictorNaniii/clean-arhitecture/tree/master/internal/model"
//	desc "github.com/VictorNaniii/clean-arhitecture/tree/master/pkg/user_v1"
//	"google.golang.org/protobuf/types/known/timestamppb"
//)
import (
	"clean-arhitecture/internal/model"
	desc "clean-arhitecture/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserInfoFromDesc(info *desc.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		First_name: info.FirstName,
		Last_name:  info.LastName,
		Age:        info.Age,
	}
}

func ToUserInfoToDesc(info *model.UserInfo) *desc.UserInfo {
	return &desc.UserInfo{
		FirstName: info.First_name,
		LastName:  info.Last_name,
		Age:       info.Age,
	}
}

func ToUserFromService(user *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt != nil {
		updatedAt = timestamppb.New(*user.UpdatedAt)
	}
	return &desc.User{
		Uuid:      user.UUID,
		Info:      ToUserInfoToDesc(&user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}
