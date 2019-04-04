package datautil

type BouyData struct {
	Year      int
	Month     int
	Day       int
	Hour      int
	Minute    int
	WVHT      float64
	SwH       float64
	SwP       float64
	WWH       float64
	WWP       float64
	SwD       string
	WWD       string
	Steepness string
	APD       float64
	MWD       int
}

type Surf interface {
	WaveSize() float64
}

func (b BouyData) WaveSize() float64 {
	return b.SwH * b.SwP
}

func (b BouyData) getWaveScore() (waveScore int) {
	waveSize := b.WaveSize()

	if waveSize >= 30 {
		return 5
	} else if waveSize >= 25 {
		return 4
	} else if waveSize >= 20 {
		return 3
	} else if waveSize >= 11 {
		return 2
	} else {
		return 1
	}
}

func (b BouyData) getSwellPeriodScore() (swellPeriodScore int) {
	if b.SwP >= 16 {
		return 5
	} else if b.SwP >= 13 {
		return 4
	} else if b.SwP >= 10 {
		return 3
	} else {
		return 1
	}
}

func (b BouyData) getWindDirectionScore() (windDirectionScore int) {
	if b.WWD == "E" {
		return 5
	} else if b.WWD == "NE" || b.WWD == "SE" {
		return 4
	} else if b.WWD == "S" {
		return 3
	} else {
		return 1
	}
}

func (b BouyData) CalculateSurfRating() (surfRating float64) {
	swellPeriodScore := float64(b.getSwellPeriodScore())
	windDirectionScore := float64(b.getWindDirectionScore())
	waveSizeScore := float64(b.getWaveScore())
	numerator := (swellPeriodScore + waveSizeScore + windDirectionScore)
	surfRating = (numerator / 3.0)
	return surfRating
}

type FetchData struct {
	url string
}

type ApiCall interface {
	Fetch() []byte
}

func (f FetchData) GetBouyData() []byte {
	var responseData []byte

	// response, err := http.Get(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer response.Body.Close()

	// responseData, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return responseData
}
