package geofence

import (
	"math"
)

func project(value float64, tileSize float64) float64 {
	return math.Floor(value / tileSize)
}

func haveIntersectingEdges(poly1 []*Point, poly2 []*Point) bool {
	for idx1 := 0; idx1 < len(poly1)-1; idx1++ {
		for idx2 := 0; idx2 < len(poly2)-1; idx2++ {
			if segmentsIntersect(poly1[idx1], poly1[idx1+1], poly2[idx2], poly2[idx2+1]) {
				return true
			}
		}
	}
	return false
}

func hasPointInPolygon(sourcePoly []*Point, targetPoly []*Point) bool {
	tPolygon := NewPolygon(targetPoly)
	for idx := 0; idx < len(sourcePoly)-1; idx++ {
		if tPolygon.Contains(sourcePoly[idx]) {
			return true
		}
	}
	return false
}

func segmentsIntersect(s1p1 *Point, s1p2 *Point, s2p1 *Point, s2p2 *Point) bool {
	// Based on http://stackoverflow.com/questions/563198/how-do-you-detect-where-two-line-segments-intersect
	p := s1p1
	r := vectorDifference(s1p2, s1p1)
	q := s2p1
	s := vectorDifference(s2p2, s2p1)

	rCrossS := vectorCrossProduct(r, s)
	qMinusP := vectorDifference(q, p)

	if rCrossS == 0 {
		if vectorCrossProduct(qMinusP, r) == 0 {
			return true
		} else {
			return false
		}
	}

	t := vectorCrossProduct(qMinusP, s) / rCrossS
	u := vectorCrossProduct(qMinusP, r) / rCrossS
	return t >= 0 && t <= 1 && u >= 0 && u <= 1
}

// here we temporarily use point struct to store vector
func vectorDifference(p1 *Point, p2 *Point) *Point {
	return NewPoint(p1.Lat()-p2.Lat(), p1.Lng()-p2.Lng())
}

func vectorCrossProduct(p1 *Point, p2 *Point) float64 {
	return p1.Lat()*p2.Lng() - p1.Lng()*p2.Lat()
}
