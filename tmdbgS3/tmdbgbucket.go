package tmdbgS3

import (
	"context"
	"errors"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//
type TmdbgBucket struct {
	clientAwsS3 *s3.Client
	iBucket     string
	iObjKey     string
}

//
func NewTmdbgBucket(strBucket, strObjKey string) (tmdbBucketObject TmdbgBucket, errorCreate error) {

	config, errConfig := config.LoadDefaultConfig(context.TODO())

	if errConfig != nil {
		return TmdbgBucket{}, errors.New("Erro ao Carregar recurso do Bucket S3")
	}

	obCliAwsS3 := s3.NewFromConfig(config)

	return TmdbgBucket{
		clientAwsS3: obCliAwsS3,
		iBucket:     strBucket,
		iObjKey:     strObjKey,
	}, nil
}

//
func (ref TmdbgBucket) GetErrorMessage() (dataStr string, iError error) {

	result, iErr := ref.clientAwsS3.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(ref.iBucket),
		Key:    aws.String(ref.iObjKey),
	})

	if iErr != nil {
		return "", errors.New("Erro ao Carregar recurso do Bucket S3")
	}

	defer result.Body.Close()

	body, errReadBody := ioutil.ReadAll(result.Body)

	if errReadBody != nil {
		return "", errors.New("Erro ao ler estrutura do objeto S3")
	}

	strReturn := string(body)

	return strReturn, nil
}
