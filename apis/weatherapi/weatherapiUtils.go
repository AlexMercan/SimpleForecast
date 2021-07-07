package weatherapi

func IsCloudy(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case Cloudy, PartlyCloudy, Overcast:
		return true
	}
	return false
}

func IsFog(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case Fog, Mist, FreezingFog:
		return true
	}
	return false
}

func IsLightRain(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case LightRain, PatchyLightRain,
		PatchyRainPossible, LightFreezingRain,
		LightRainShower, PatchyLightDrizzle,
		LightDrizzle, FreezingDrizzle,
		HeavyFreezingDrizzle, ModerateOrHeavyRainShower,
		PatchyFreezingDrizzlePossible:
		return true
	}
	return false
}

func IsHeavyRain(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case HeavyRain, HeavyRainAtTimes, ModerateRain,
		ModerateRainAtTimes, TorrentialRainShower,
		ModerateOrHeavyFreezingRain:
		return true
	}
	return false
}

func IsLightSnow(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case LightSnow, LightSnowShowers,
		PatchyLightSnow, PatchySnowPossible,
		LightShowersOfIcePellets, BlowingSnow,
		IcePellets:
		return true
	}
	return false
}

func IsHeavySnow(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case HeavySnow, ModerateOrHeavySnowShowers,
		PatchyModerateSnow, ModerateSnow,
		PatchyHeavySnow, ModerateOrHeavyShowersOfIcePellets,
		Blizzard:
		return true
	}
	return false
}

func IsSleet(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case LightSleet, LightSleetShowers,
		ModerateOrHeavySleetShowers, PatchySleetPossible,
		ModerateOrHeavySleet:
		return true
	}
	return false
}

func IsSunny(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case Sunny:
		return true
	}
	return false
}

func IsHeavyRainThunder(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case ModerateOrHeavyRainWithThunder, ThunderyOutbreaksPossible:
		return true
	}
	return false
}

func IsHeavyShowersThunder(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case PatchyLightRainWithThunder:
		return true
	}
	return false
}

func IsSnowShowerThunder(weatherConditionCode int) bool {
	switch weatherConditionCode {
	case PatchyLightSnowWithThunder, ModerateOrHeavySnowWithThunder:
		return true
	}
	return false
}

