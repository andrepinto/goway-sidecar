package collector

import (
	"google.golang.org/api/support/bundler"
	"time"
	"fmt"
	"github.com/twinj/uuid"
	"github.com/jehiah/go-strftime"
	"github.com/andrepinto/goway-sidecar/outputs"
	"strconv"
	log "github.com/sirupsen/logrus"
)

type HttpLoggerAction struct {
	htppLoggerBundler *bundler.Bundler
	Output outputs.Output
}


func(tr *HttpLoggerAction) Fire(httpLg *HttpLogger) (string, error){


	id:=uuid.NewV4().String()

	httpLg.Base.Type = HTTP_LOGGER_ACTION
	httpLg.Id= id
	httpLg.Timestamp = GetTimestamp(time.Now())
	httpLg.Key = httpLg.Key


	httpLg.callback = func(lg *HttpLogger){
		fmt.Println(fmt.Sprintf("callback %s", lg.Id))
	}
	if err := tr.htppLoggerBundler.Add(httpLg, 0); err != nil {
		panic(err)
	}

	return id, nil
}


func GetTimestamp(t time.Time) string {
	return strftime.Format("%Y-%m-%dT%H:%M:%S%z", t)
}


func (tc *HttpLoggerAction) Close() {
	tc.htppLoggerBundler.Flush()
}

func(tc *HttpLoggerAction) sendHttpLogger(logs []*HttpLogger) (error){
	arr := []*outputs.HttpLoggerClient{}
	for _, item := range logs {

		date, err := DateTimeToDate(item.Data.Time.Year, item.Data.Time.Month, item.Data.Time.Day, item.Data.Time.Hour, item.Data.Time.Min, item.Data.Time.Sec)
		if err!=nil{
			return err
		}

		tr := &outputs.HttpLoggerClient{
			RequestId: item.Data.RequestId,
			Base: outputs.BaseClientRequest{
				Properties: item.Properties,
				Id: item.Key,
				Key: item.Key,
				Context: item.Context,
			},
			Data: outputs.HttpLoggerRequestClient{
				Protocol:item.Data.Protocol,
				Uri: item.Data.Uri,
				Host: item.Data.Host,
				Ip: item.Data.Ip,
				ElapsedTime: item.Data.ElapsedTime,
				BasePath: item.Data.BasePath,
				RequestHeader: item.Data.RequestHeader,
				RequestBody: item.Data.RequestBody,
				Version: item.Data.Version,
				Method: item.Data.Method,
				Time: date,
				ResponseBody: item.Data.ResponseBody,
				Tags: item.Data.Tags,
				Status: item.Data.Status,
				ServicePath: item.Data.ServicePath,
				Metadata: item.Data.Metadata,
			},
			Timestamp: item.Timestamp,
		}
		arr = append(arr, tr)
	}

	log.Debug("sending...")

	tc.Output.Send(arr)

	log.Debug("ok.")

	return nil
}

func DateTimeToDate(year string, month string, day string, hour string, min string, sec string) (time.Time, error){
	iy, err :=strconv.Atoi(year)
	im, err :=strconv.Atoi(month)
	id, err :=strconv.Atoi(day)
	ih, err :=strconv.Atoi(hour)
	imin, err :=strconv.Atoi(min)
	is, err :=strconv.Atoi(sec)

	if err!=nil{
		return time.Now(), err
	}

	date := time.Date(iy, time.Month(im), id, ih, imin, is, 0, time.UTC)
	return date, nil
}