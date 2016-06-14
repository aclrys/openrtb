package openrtb

// 5.15 QAG Media Ratings
//
// The following table lists the media ratings used in describing content based on the QAG categorization.
// Refer to http://www.iab.net/ne_guidelines for more information.
type QAGMediaRating int8

const (
	QAGMediaRatingAllAudiences    QAGMediaRating = 1 // 1 All Audiences
	QAGMediaRatingEveryoneOver12  QAGMediaRating = 2 // 2 Everyone Over 12
	QAGMediaRatingMatureAudiences QAGMediaRating = 3 // 3 Mature Audiences
)
