package version

import (
	"fmt"
	"strconv"
	"strings"
)

type Version uint64

func MustParse(s string) Version {
	v, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return v
}

func Parse(s string) (v Version, err error) {
	pieces := strings.SplitN(s, ".", 3)

	if len(pieces) != 3 {
		return v, ErrInsufficientParts
	}

	major, err := strconv.ParseUint(pieces[0], 10, 16)
	if err != nil {
		return v, err
	}

	minor, err := strconv.ParseUint(pieces[1], 10, 16)
	if err != nil {
		return v, err
	}

	patch, err := strconv.ParseUint(pieces[2], 10, 16)
	if err != nil {
		return v, err
	}

	v = Version(patch | minor<<16 | major<<32)

	return v, nil
}

func (v *Version) UnmarshalText(text []byte) error {
	parsed, err := Parse(string(text))
	if err != nil {
		return err
	}

	*v = parsed
	return nil
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", uint16(v>>32), uint16(v>>16), uint16(v))
}

//type Version struct {
//	Major uint16
//	Minor uint16
//	Patch uint16
//}
//
//func Parse(s string) (v Version, err error) {
//	pieces := strings.SplitN(s, ".", 3)
//
//	if len(pieces) != 3 {
//		return v, ErrInsufficientParts
//	}
//
//	major, err := strconv.ParseUint(pieces[0], 10, 16)
//	if err != nil {
//		return v, err
//	}
//
//	minor, err := strconv.ParseUint(pieces[1], 10, 16)
//	if err != nil {
//		return v, err
//	}
//
//	patch, err := strconv.ParseUint(pieces[2], 10, 16)
//	if err != nil {
//		return v, err
//	}
//
//	v.Major = uint16(major)
//	v.Minor = uint16(minor)
//	v.Patch = uint16(patch)
//	return v, nil
//}
//
//func (v *Version) UnmarshalText(text []byte) error {
//	parsed, err := Parse(string(text))
//	if err != nil {
//		return err
//	}
//
//	v.Major = parsed.Major
//	v.Minor = parsed.Minor
//	v.Patch = parsed.Patch
//	return nil
//}
//
//func (v Version) String() string {
//	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
//}
