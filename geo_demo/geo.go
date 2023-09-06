package main

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geo"
	"github.com/paulmach/orb/geojson"
	"github.com/paulmach/orb/planar"
	"github.com/paulmach/orb/resample"
	"golang.org/x/xerrors"
	"log"
	"sync"
)

type Coordinates struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

func ExtractCircleEdgePoints(coordinates []Coordinates) ([]HLPoint, error) {
	fc := geojson.NewFeatureCollection()
	points := []orb.Point{}
	for _, point := range coordinates {
		points = append(points, [2]float64{point.Lng, point.Lat})
		fc.Append(geojson.NewFeature(orb.MultiPolygon{orb.Polygon{points}}))
	}
	geoPoints, err := GetPointsInside(fc, 600)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return geoPoints, nil
}

func GetPointsInside(fc *geojson.FeatureCollection, distance float64) ([]HLPoint, error) {
	var mp = orb.MultiPolygon{}
	mp = MPofFC(fc)
	points := orb.MultiPoint{}
	pointsLock := new(sync.Mutex)
	bound := mp.Bound()
	minX, minY, maxX, maxY := bound.Min.X(), bound.Min.Y(), bound.Max.X(), bound.Max.Y()
	yLineString := orb.LineString{orb.Point{minX, minY}, orb.Point{minX, maxY}}
	yLineString = resample.ToInterval(yLineString, geo.DistanceHaversine, distance)
	wg := sync.WaitGroup{}
	for i := 0; i < len(yLineString); i++ {
		wg.Add(1)
		go func(i int) {
			xLineString := orb.LineString{orb.Point{minX, yLineString[i].Y()}, orb.Point{maxX, yLineString[i].Y()}}
			xLineString = resample.ToInterval(xLineString, geo.DistanceHaversine, distance)
			for _, p := range xLineString {
				if planar.MultiPolygonContains(mp, p) {
					pointsLock.Lock()
					points = append(points, p)
					pointsLock.Unlock()
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	jpoints := []HLPoint{}
	for _, point := range points {
		jpoints = append(jpoints, HLPoint{Lat: point.Y(), Lng: point.X()})
	}
	return jpoints, nil
}

type HLPoint struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func MPofFC(fc *geojson.FeatureCollection) orb.MultiPolygon {
	mp := orb.MultiPolygon{}
	for _, f := range fc.Features {
		switch metry := f.Geometry.(type) {
		case orb.MultiPolygon:
			mp = append(mp, metry...)
		case orb.Polygon:
			mp = append(mp, metry)
		default:
			log.Println("geometry type:", f.Geometry.GeoJSONType())
		}
	}
	return mp
}

var coordinatesJson = `[
    {
      "lng": 112.98311,
      "lat": 28.24254
    },
    {
      "lng": 112.989003,
      "lat": 28.239787
    },
    {
      "lng": 112.990004,
      "lat": 28.243749
    },
    {
      "lng": 112.989717,
      "lat": 28.246028
    },
    {
      "lng": 112.990836,
      "lat": 28.246314
    },
    {
      "lng": 112.991181,
      "lat": 28.244386
    },
    {
      "lng": 112.9913,
      "lat": 28.243659
    },
    {
      "lng": 112.990453,
      "lat": 28.239633
    },
    {
      "lng": 112.992076,
      "lat": 28.239165
    },
    {
      "lng": 112.996321,
      "lat": 28.239555
    },
    {
      "lng": 113.007003,
      "lat": 28.241835
    },
    {
      "lng": 113.00702649443883,
      "lat": 28.245629817756605
    },
    {
      "lng": 113.00784466282926,
      "lat": 28.24560063522476
    },
    {
      "lng": 113.007685,
      "lat": 28.242171
    },
    {
      "lng": 113.01115,
      "lat": 28.242412
    },
    {
      "lng": 113.019518,
      "lat": 28.240673
    },
    {
      "lng": 113.019955,
      "lat": 28.242452
    },
    {
      "lng": 113.019963,
      "lat": 28.243475
    },
    {
      "lng": 113.020631,
      "lat": 28.243482
    },
    {
      "lng": 113.020514,
      "lat": 28.240466
    },
    {
      "lng": 113.024651,
      "lat": 28.239539
    },
    {
      "lng": 113.025398,
      "lat": 28.238867
    },
    {
      "lng": 113.025751,
      "lat": 28.237577
    },
    {
      "lng": 113.025671,
      "lat": 28.235826
    },
    {
      "lng": 113.025351,
      "lat": 28.234355
    },
    {
      "lng": 113.028368,
      "lat": 28.234357
    },
    {
      "lng": 113.028417,
      "lat": 28.233782
    },
    {
      "lng": 113.025182,
      "lat": 28.23366
    },
    {
      "lng": 113.024432,
      "lat": 28.232109
    },
    {
      "lng": 113.022715,
      "lat": 28.22994
    },
    {
      "lng": 113.019433,
      "lat": 28.227166
    },
    {
      "lng": 113.019862,
      "lat": 28.224003
    },
    {
      "lng": 113.020248,
      "lat": 28.222243
    },
    {
      "lng": 113.024968,
      "lat": 28.222732
    },
    {
      "lng": 113.028401,
      "lat": 28.222553
    },
    {
      "lng": 113.032821,
      "lat": 28.221087
    },
    {
      "lng": 113.032307,
      "lat": 28.216739
    },
    {
      "lng": 113.039474,
      "lat": 28.214432
    },
    {
      "lng": 113.044066,
      "lat": 28.199341
    },
    {
      "lng": 113.047081,
      "lat": 28.200344
    },
    {
      "lng": 113.047301,
      "lat": 28.199143
    },
    {
      "lng": 113.044407,
      "lat": 28.198448
    },
    {
      "lng": 113.050353,
      "lat": 28.189848
    },
    {
      "lng": 113.05762608677878,
      "lat": 28.192480453868573
    },
    {
      "lng": 113.0595288637403,
      "lat": 28.19055187946097
    },
    {
      "lng": 113.054063,
      "lat": 28.187625
    },
    {
      "lng": 113.057944,
      "lat": 28.18525
    },
    {
      "lng": 113.059442,
      "lat": 28.184359
    },
    {
      "lng": 113.053473,
      "lat": 28.174313
    },
    {
      "lng": 113.052205,
      "lat": 28.171371
    },
    {
      "lng": 113.055648,
      "lat": 28.169975
    },
    {
      "lng": 113.05415,
      "lat": 28.167083
    },
    {
      "lng": 113.052887,
      "lat": 28.165713
    },
    {
      "lng": 113.051022,
      "lat": 28.164872
    },
    {
      "lng": 113.051057,
      "lat": 28.162834
    },
    {
      "lng": 113.052465,
      "lat": 28.157013
    },
    {
      "lng": 113.049203,
      "lat": 28.157007
    },
    {
      "lng": 113.048173,
      "lat": 28.154129
    },
    {
      "lng": 113.045212,
      "lat": 28.15076
    },
    {
      "lng": 113.044547,
      "lat": 28.148583
    },
    {
      "lng": 113.04577,
      "lat": 28.130361
    },
    {
      "lng": 113.049698,
      "lat": 28.118769
    },
    {
      "lng": 113.029984,
      "lat": 28.109679
    },
    {
      "lng": 113.016862,
      "lat": 28.108989
    },
    {
      "lng": 113.019047,
      "lat": 28.10054
    },
    {
      "lng": 113.016855,
      "lat": 28.093074
    },
    {
      "lng": 113.010465,
      "lat": 28.094128
    },
    {
      "lng": 113.010426,
      "lat": 28.098665
    },
    {
      "lng": 112.999706,
      "lat": 28.097366
    },
    {
      "lng": 112.9792507786874,
      "lat": 28.096790025460734
    },
    {
      "lng": 112.97779501341881,
      "lat": 28.105352137491433
    },
    {
      "lng": 112.97747508645307,
      "lat": 28.113232340418115
    },
    {
      "lng": 112.978875,
      "lat": 28.118989
    },
    {
      "lng": 112.978979,
      "lat": 28.120934
    },
    {
      "lng": 112.979231,
      "lat": 28.123766
    },
    {
      "lng": 112.978974,
      "lat": 28.129443
    },
    {
      "lng": 112.975165,
      "lat": 28.131298
    },
    {
      "lng": 112.969833,
      "lat": 28.13495
    },
    {
      "lng": 112.970412,
      "lat": 28.136085
    },
    {
      "lng": 112.971335,
      "lat": 28.139984
    },
    {
      "lng": 112.967409,
      "lat": 28.138205
    },
    {
      "lng": 112.960284,
      "lat": 28.137466
    },
    {
      "lng": 112.958782,
      "lat": 28.137466
    },
    {
      "lng": 112.950275,
      "lat": 28.13524
    },
    {
      "lng": 112.943313,
      "lat": 28.137252
    },
    {
      "lng": 112.94309,
      "lat": 28.146076
    },
    {
      "lng": 112.950591,
      "lat": 28.151418
    },
    {
      "lng": 112.953128,
      "lat": 28.161584
    },
    {
      "lng": 112.952023,
      "lat": 28.17117
    },
    {
      "lng": 112.952817,
      "lat": 28.182093
    },
    {
      "lng": 112.955498,
      "lat": 28.1951
    },
    {
      "lng": 112.952129,
      "lat": 28.19517
    },
    {
      "lng": 112.949166,
      "lat": 28.19811
    },
    {
      "lng": 112.94058,
      "lat": 28.201569
    },
    {
      "lng": 112.936624,
      "lat": 28.201892
    },
    {
      "lng": 112.932744,
      "lat": 28.202461
    },
    {
      "lng": 112.931053,
      "lat": 28.203144
    },
    {
      "lng": 112.930113,
      "lat": 28.200896
    },
    {
      "lng": 112.928104,
      "lat": 28.199255
    },
    {
      "lng": 112.928091,
      "lat": 28.1975
    },
    {
      "lng": 112.923945,
      "lat": 28.19888
    },
    {
      "lng": 112.921816,
      "lat": 28.197631
    },
    {
      "lng": 112.916114,
      "lat": 28.198486
    },
    {
      "lng": 112.912172,
      "lat": 28.197449
    },
    {
      "lng": 112.887808,
      "lat": 28.182298
    },
    {
      "lng": 112.875666,
      "lat": 28.182369
    },
    {
      "lng": 112.870486,
      "lat": 28.181833
    },
    {
      "lng": 112.865736,
      "lat": 28.185079
    },
    {
      "lng": 112.863002,
      "lat": 28.187833
    },
    {
      "lng": 112.860869,
      "lat": 28.191269
    },
    {
      "lng": 112.856498,
      "lat": 28.203804
    },
    {
      "lng": 112.851075,
      "lat": 28.221435
    },
    {
      "lng": 112.871303,
      "lat": 28.221768
    },
    {
      "lng": 112.889986,
      "lat": 28.220211
    },
    {
      "lng": 112.890787,
      "lat": 28.221709
    },
    {
      "lng": 112.900672,
      "lat": 28.22187
    },
    {
      "lng": 112.907674,
      "lat": 28.222158
    },
    {
      "lng": 112.914097,
      "lat": 28.222464
    },
    {
      "lng": 112.91889,
      "lat": 28.229954
    },
    {
      "lng": 112.924185,
      "lat": 28.245272
    },
    {
      "lng": 112.930784,
      "lat": 28.251181
    },
    {
      "lng": 112.936839,
      "lat": 28.256033
    },
    {
      "lng": 112.937195,
      "lat": 28.256143
    },
    {
      "lng": 112.941092,
      "lat": 28.258638
    },
    {
      "lng": 112.944989,
      "lat": 28.260567
    },
    {
      "lng": 112.947599,
      "lat": 28.263667
    },
    {
      "lng": 112.948492,
      "lat": 28.266993
    },
    {
      "lng": 112.957971,
      "lat": 28.26036
    },
    {
      "lng": 112.962637,
      "lat": 28.254659
    },
    {
      "lng": 112.98045689503567,
      "lat": 28.25518240858986
    },
    {
      "lng": 112.98124298440018,
      "lat": 28.2520925476419
    },
    {
      "lng": 112.961686,
      "lat": 28.251496
    },
    {
      "lng": 112.963601,
      "lat": 28.247291
    },
    {
      "lng": 112.963748,
      "lat": 28.24442
    },
    {
      "lng": 112.962602,
      "lat": 28.240779
    },
    {
      "lng": 112.965919,
      "lat": 28.227383
    },
    {
      "lng": 112.973937,
      "lat": 28.227496
    },
    {
      "lng": 112.978191,
      "lat": 28.24555
    }
  ]`
