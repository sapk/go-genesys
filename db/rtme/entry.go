package rtme

import (
	"strings"
	"time"

	"github.com/sapk/go-genesys/db"
)

func parseStartEndOfDay(day string) (time.Time, time.Time, error) {
	start, err := time.Parse("2006-01-02", day)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	start = start.UTC() //Reset to UTC
	end := start.Add(24 * time.Hour)
	return start, end, nil
}

//LoginEntry a db LOGIN entry
type LoginEntry struct {
	SWITCHDBID int    `xorm:"'SWITCHDBID'"`
	DNDBID     int    `xorm:"'DNDBID'"`
	QUEUEDBID  int    `xorm:"'QUEUEDBID'"`
	AGENTDBID  int    `xorm:"'AGENTDBID'"`
	PLACEDBID  int    `xorm:"'PLACEDBID'"`
	STATUS     int    `xorm:"'STATUS'"`
	TIME       int    `xorm:"'TIME'"`
	LOGINID    string `xorm:"varchar(255) null 'LOGINID'"`
}

//TableName in database
func (*LoginEntry) TableName() string {
	return "LOGIN"
}

//GetLoginEntriesOfDay list the entries of selected day
func GetLoginEntriesOfDay(d *db.DB, day string) ([]LoginEntry, error) {
	start, end, err := parseStartEndOfDay(day)
	if err != nil {
		return nil, err
	}
	var entries []LoginEntry
	err = d.Engine.Where("LOGINID != '' AND TIME BETWEEN ? AND ?", start.Unix(), end.Unix()).OrderBy("TIME ASC").Find(&entries)
	for _, et := range entries {
		//Clean LOGINID
		et.LOGINID = strings.TrimSpace(et.LOGINID)
	}
	return entries, err
}

//	results, err := engine.Query(fmt.Sprintf("USE RTME; SELECT TIME, PLACEDBID, AGENTDBID, LOGINID, STATUS FROM dbo.LOGIN WHERE LOGINID != '' AND TIME BETWEEN %d AND %d  ORDER BY TIME DESC", start.Unix(), end.Unix()))

//	results, err := engine.Query(fmt.Sprintf("USE RTME; SELECT AgentDBID, PlaceDBID, ConnID, Status, StartTime, EndTime FROM dbo.STATUS WHERE Status != '23' AND AgentDBID != '' AND (StartTime BETWEEN %d AND %d OR EndTime BETWEEN %d AND %d) ORDER BY StartTime ASC", start.Unix(), end.Unix(), start.Unix(), end.Unix()))

//QueueEntry a db QINFO entry
type QueueEntry struct {
	QueueDBID int `xorm:"'QueueDBID'"`
	ConnID    int `xorm:"'ConnID'"`
	Status    int `xorm:"'Status'"`
	StartTime int `xorm:"'StartTime'"`
	Duration  int `xorm:"'Duration'"`
	EndTime   int `xorm:"'EndTime'"`
}

//TableName in database
func (*QueueEntry) TableName() string {
	return "QINFO"
}

//GetQInfoEntriesOfDay list the entries of selected day
func GetQInfoEntriesOfDay(d *db.DB, day string) ([]QueueEntry, error) {
	start, end, err := parseStartEndOfDay(day)
	if err != nil {
		return nil, err
	}
	var entries []QueueEntry
	err = d.Engine.Where("StartTime BETWEEN ? AND ? OR EndTime BETWEEN ? AND ? OR (StartTime < ? AND EndTime > ?)", start.Unix(), end.Unix(), start.Unix(), end.Unix(), start.Unix(), end.Unix()).OrderBy("StartTime ASC").Find(&entries)
	return entries, err
}

//StatusEntry a db STATUS entry
type StatusEntry struct {
	ID             int    `xorm:"'ID'"`
	AgentDBID      int    `xorm:"'AgentDBID'"`
	PlaceDBID      int    `xorm:"'PlaceDBID'"`
	Status         int    `xorm:"'Status'"`
	StartTime      int    `xorm:"'StartTime'"`
	Duration       int    `xorm:"'Duration'"`
	EndTime        int    `xorm:"'EndTime'"`
	ConnID         int    `xorm:"decimal(20,0) 'ConnID'"`
	StartLocalTime string `xorm:"varchar(50) null 'StartLocalTime'"`
	EndLocalTime   string `xorm:"varchar(50) null 'EndLocalTime'"`
}

//TableName in database
func (*StatusEntry) TableName() string {
	return "STATUS"
}

//GetStatusEntriesOfDay list the entries of selected day
func GetStatusEntriesOfDay(d *db.DB, day string) ([]StatusEntry, error) {
	start, end, err := parseStartEndOfDay(day)
	if err != nil {
		return nil, err
	}
	var entries []StatusEntry
	err = d.Engine.Where("StartTime BETWEEN ? AND ? OR EndTime BETWEEN ? AND ? OR (StartTime < ? AND EndTime > ?)", start.Unix(), end.Unix(), start.Unix(), end.Unix(), start.Unix(), end.Unix()).OrderBy("StartTime ASC").Find(&entries)
	return entries, err
}

//GraphEntry a graph table entry (calculated from db entry)
type GraphEntry struct {
	Agent    string
	Sessions []Session
}

//Event a user event
type Event struct {
	State int
	Place int
	Time  int64
}

//Session a login session
type Session struct {
	Start int64
	Place int
	End   int64
}

//LoginToEventByUser regroup by user
func LoginToEventByUser(loginEntries []LoginEntry) (map[string][]Event, map[string]int) {
	byUser := make(map[string][]Event)
	userIDList := make(map[string]int)
	for _, e := range loginEntries {
		if e.LOGINID == "" {
			continue //Skip undefined
		}
		if _, ok := byUser[e.LOGINID]; !ok {
			byUser[e.LOGINID] = make([]Event, 0) //init user
			userIDList[e.LOGINID] = e.AGENTDBID  //Save dbids
		}
		byUser[e.LOGINID] = append(byUser[e.LOGINID], Event{
			State: e.STATUS,
			Place: e.PLACEDBID,
			Time:  int64(e.TIME),
		})
	}
	return byUser, userIDList
}

//LoginEventByUserToSession regroup by user
func LoginEventByUserToSession(start, end time.Time, loginEventsByUser map[string][]Event) map[string][]Session {
	bySession := make(map[string][]Session)
	for user, events := range loginEventsByUser {
		var endSession int64
		var endSessionPlace int
		bySession[user] = make([]Session, 0) //init user
		for _, evt := range events {
			if evt.State == 1 && endSession == 0 { //unclosed session of day
				bySession[user] = append(bySession[user], Session{
					Start: evt.Time,
					Place: evt.Place,
					End:   end.Unix(),
				})
				//Reset
				endSession = 0
				//startSession = ""
				continue
			}
			if evt.State == 0 && endSession == 0 { //catch logout
				endSession = evt.Time
				endSessionPlace = evt.Place
				continue
			}
			if evt.State == 1 && endSession != 0 { //catch login
				bySession[user] = append(bySession[user], Session{
					Start: evt.Time,
					Place: evt.Place, //TODO consolidate via place
					End:   endSession,
				})
				//Reset
				endSession = 0
				//startSession = ""
				continue
			}
			//log.Warn().Msgf("Uncatch event of user %s: %v", user, evt)
		}
		if endSession != 0 {
			//We have a session opened the day before
			bySession[user] = append(bySession[user], Session{
				Start: start.Unix(),
				Place: endSessionPlace,
				End:   endSession,
			})
		}
	}
	return bySession
	/*
	 */
}

//FormattedLoginResp format the login to be used by api
type FormattedLoginResp struct {
	Start    int64
	End      int64
	Sessions map[string][]Session
	Users    map[string]int
}

//FormattedLoginEntriesOfDay formatted login entry
func FormattedLoginEntriesOfDay(d *db.DB, day string) (*FormattedLoginResp, error) {
	et, err := GetLoginEntriesOfDay(d, day)
	if err != nil {
		return nil, err
	}
	start, end, err := parseStartEndOfDay(day)
	if err != nil {
		return nil, err
	}

	events, userList := LoginToEventByUser(et)
	return &FormattedLoginResp{
		Start:    start.Unix(),
		End:      end.Unix(),
		Sessions: LoginEventByUserToSession(start, end, events),
		Users:    userList,
	}, nil
}

//GetGraphEntriesOfDay calculate graph entry
func GetGraphEntriesOfDay(d *db.DB, day string) ([]GraphEntry, error) {
	et, err := GetLoginEntriesOfDay(d, day)
	if err != nil {
		return nil, err
	}
	start, end, err := parseStartEndOfDay(day)
	if err != nil {
		return nil, err
	}
	events, _ := LoginToEventByUser(et)
	bySession := LoginEventByUserToSession(start, end, events)
	//Reformat
	i := 0
	response := make([]GraphEntry, len(bySession))
	for user, sessions := range bySession {
		response[i] = GraphEntry{
			Agent:    user,
			Sessions: sessions,
		}
		i++
	}
	return response, nil
}
