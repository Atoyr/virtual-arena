// internal/storage/s3.go
package storage

// import (
//   "bytes"
//   "fmt"
//   "io"
//   "time"
//
//   "github.com/aws/aws-sdk-go/aws"
//   "github.com/aws/aws-sdk-go/service/s3"
// )
//
// type S3Storage struct {
//   client   *s3.S3
//   bucket   string
//   cdnURL   string // CloudFront ドメイン
// }
//
// func NewS3Storage(client *s3.S3, bucket, cdnURL string) *S3Storage {
//   return &S3Storage{client, bucket, cdnURL}
// }
//
// func (s *S3Storage) SaveMapJSON(tenantID, mapID string, data []byte) error {
//   key := fmt.Sprintf("%s/maps/%s/map.json", tenantID, mapID)
//   _, err := s.client.PutObject(&s3.PutObjectInput{
//     Bucket: aws.String(s.bucket),
//     Key:    aws.String(key),
//     Body:   bytes.NewReader(data),
//     ACL:    aws.String("private"),
//   })
//   return err
// }
//
// func (s *S3Storage) LoadMapJSON(tenantID, mapID string) ([]byte, error) {
//   key := fmt.Sprintf("%s/maps/%s/map.json", tenantID, mapID)
//   out, err := s.client.GetObject(&s3.GetObjectInput{
//     Bucket: aws.String(s.bucket), Key: aws.String(key),
//   })
//   if err != nil {
//     return nil, err
//   }
//   defer out.Body.Close()
//   buf := new(bytes.Buffer)
//   _, err = io.Copy(buf, out.Body)
//   return buf.Bytes(), err
// }
//
// func (s *S3Storage) SaveTile(tenantID, mapID, z, x, y string, data io.Reader) error {
//   key := fmt.Sprintf("%s/maps/%s/tiles/%s/%s/%s.png", tenantID, mapID, z, x, y)
//   _, err := s.client.PutObject(&s3.PutObjectInput{
//     Bucket: aws.String(s.bucket),
//     Key:    aws.String(key),
//     Body:   data,
//     ACL:    aws.String("public-read"), // 公開 or Signed URL
//   })
//   return err
// }
//
// func (s *S3Storage) TileURL(tenantID, mapID, z, x, y string) (string, error) {
//   // CDN あり
//   return fmt.Sprintf("%s/%s/maps/%s/tiles/%s/%s/%s.png",
//     s.cdnURL, tenantID, mapID, z, x, y), nil
// }
//
