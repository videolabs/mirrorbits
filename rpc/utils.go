// Copyright (c) 2014-2019 Ludovic Fauvet
// Licensed under the MIT license

package rpc

import (
	"github.com/etix/mirrorbits/mirrors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MirrorToRPC(m *mirrors.Mirror) (*Mirror, error) {
	stateSince := timestamppb.New(m.StateSince.Time)
	lastSync := timestamppb.New(m.LastSync.Time)
	lastSuccessfulSync := timestamppb.New(m.LastSuccessfulSync.Time)
	lastModTime := timestamppb.New(m.LastModTime.Time)

	return &Mirror{
		ID:                   int32(m.ID),
		Name:                 m.Name,
		HttpURL:              m.HttpURL,
		RsyncURL:             m.RsyncURL,
		FtpURL:               m.FtpURL,
		SponsorName:          m.SponsorName,
		SponsorURL:           m.SponsorURL,
		SponsorLogoURL:       m.SponsorLogoURL,
		AdminName:            m.AdminName,
		AdminEmail:           m.AdminEmail,
		CustomData:           m.CustomData,
		ContinentOnly:        m.ContinentOnly,
		CountryOnly:          m.CountryOnly,
		ASOnly:               m.ASOnly,
		Score:                int32(m.Score),
		Latitude:             m.Latitude,
		Longitude:            m.Longitude,
		ContinentCode:        m.ContinentCode,
		CountryCodes:         m.CountryCodes,
		ExcludedCountryCodes: m.ExcludedCountryCodes,
		Asnum:                uint32(m.Asnum),
		Comment:              m.Comment,
		Enabled:              m.Enabled,
		HttpUp:               m.HttpUp,
		HttpsUp:              m.HttpsUp,
		HttpDownReason:       m.HttpDownReason,
		HttpsDownReason:      m.HttpsDownReason,
		StateSince:           stateSince,
		AllowRedirects:       int32(m.AllowRedirects),
		LastSync:             lastSync,
		LastSuccessfulSync:   lastSuccessfulSync,
		LastModTime:          lastModTime,
	}, nil
}

func MirrorFromRPC(m *Mirror) (*mirrors.Mirror, error) {
	var err error

	if err = m.StateSince.CheckValid(); err != nil {
		return nil, err
	}
	stateSince := m.StateSince.AsTime()

	if err = m.LastSync.CheckValid(); err != nil {
		return nil, err
	}
	lastSync := m.LastSync.AsTime()

	if err = m.LastSuccessfulSync.CheckValid(); err != nil {
		return nil, err
	}
	lastSuccessfulSync := m.LastSuccessfulSync.AsTime()

	if err = m.LastModTime.CheckValid(); err != nil {
		return nil, err
	}
	lastModTime := m.LastModTime.AsTime()

	return &mirrors.Mirror{
		ID:                   int(m.ID),
		Name:                 m.Name,
		HttpURL:              m.HttpURL,
		RsyncURL:             m.RsyncURL,
		FtpURL:               m.FtpURL,
		SponsorName:          m.SponsorName,
		SponsorURL:           m.SponsorURL,
		SponsorLogoURL:       m.SponsorLogoURL,
		AdminName:            m.AdminName,
		AdminEmail:           m.AdminEmail,
		CustomData:           m.CustomData,
		ContinentOnly:        m.ContinentOnly,
		CountryOnly:          m.CountryOnly,
		ASOnly:               m.ASOnly,
		Score:                int(m.Score),
		Latitude:             m.Latitude,
		Longitude:            m.Longitude,
		ContinentCode:        m.ContinentCode,
		CountryCodes:         m.CountryCodes,
		ExcludedCountryCodes: m.ExcludedCountryCodes,
		Asnum:                uint(m.Asnum),
		Comment:              m.Comment,
		Enabled:              m.Enabled,
		HttpUp:               m.HttpUp,
		HttpsUp:              m.HttpsUp,
		HttpDownReason:       m.HttpDownReason,
		HttpsDownReason:      m.HttpsDownReason,
		StateSince:           mirrors.Time{}.FromTime(stateSince),
		AllowRedirects:       mirrors.Redirects(m.AllowRedirects),
		LastSync:             mirrors.Time{}.FromTime(lastSync),
		LastSuccessfulSync:   mirrors.Time{}.FromTime(lastSuccessfulSync),
		LastModTime:          mirrors.Time{}.FromTime(lastModTime),
	}, nil
}
