package snowflake

import "strconv"

func ConvertToSnowflake(input string) (ID, error) {
	parsed, err := strconv.ParseInt(input, 10, 64)
	return ID(parsed), err
}

func ConvertFromSnowflake(input ID) string {
	return strconv.FormatInt(int64(input), 10)
}

func ConvertToStrings(snowflakes []ID) []string {
	result := make([]string, len(snowflakes))

	for i, v := range snowflakes {
		result[i] = v.String()
	}

	return result
}
