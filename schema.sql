CREATE TABLE calendar (
	service_id VARCHAR(14) NOT NULL,
	monday INTEGER NOT NULL,
	tuesday INTEGER NOT NULL,
	wednesday INTEGER NOT NULL,
	thursday INTEGER NOT NULL,
	friday INTEGER NOT NULL,
	saturday INTEGER NOT NULL,
	sunday INTEGER NOT NULL,
	start_date INTEGER NOT NULL,
	end_date INTEGER NOT NULL
);
CREATE INDEX idx_calendar_service_id ON calendar (service_id);
CREATE TABLE calendar_dates (
	service_id VARCHAR(14) NOT NULL,
	date INTEGER NOT NULL,
	exception_type INTEGER NOT NULL
);
CREATE INDEX idx_calendar_dates_service_id ON calendar_dates (service_id);
CREATE TABLE routes (
	route_id VARCHAR(14) NOT NULL,
	agency_id VARCHAR(6) NOT NULL,
	route_short_name VARCHAR(6),
	route_long_name VARCHAR(100) NOT NULL,
	route_desc VARCHAR(28) NOT NULL,
	route_type INTEGER NOT NULL,
	route_color VARCHAR(6) NOT NULL,
	route_text_color VARCHAR(6) NOT NULL
);
CREATE INDEX idx_routes_route_id ON routes (route_id);
CREATE TABLE stop_times (
	trip_id VARCHAR(28) NOT NULL,
	arrival_time VARCHAR(8),
	departure_time VARCHAR(8),
	stop_id VARCHAR(7) NOT NULL,
	stop_sequence INTEGER NOT NULL,
	stop_headsign VARCHAR(35),
	pickup_type INTEGER NOT NULL,
	drop_off_type INTEGER NOT NULL,
	shape_dist_traveled FLOAT
);
CREATE INDEX idx_stop_times_trip_id ON stop_times (trip_id);
CREATE TABLE stops (
	stop_id VARCHAR(7) NOT NULL,
	stop_code INTEGER,
	stop_name VARCHAR(79) NOT NULL,
	stop_lat FLOAT NOT NULL,
	stop_lon FLOAT NOT NULL,
	location_type INTEGER,
	parent_station VARCHAR(7),
	wheelchair_boarding INTEGER NOT NULL,
	platform_code VARCHAR(6)
);
CREATE INDEX idx_stops_stop_id ON stops (stop_id);
CREATE TABLE trips (
	route_id VARCHAR(14) NOT NULL,
	service_id VARCHAR(14) NOT NULL,
	trip_id VARCHAR(28) NOT NULL,
	shape_id VARCHAR(19) NOT NULL,
	trip_headsign VARCHAR(94) NOT NULL,
	direction_id INTEGER NOT NULL,
	block_id VARCHAR(6),
	wheelchair_accessible INTEGER NOT NULL
);
CREATE INDEX idx_trips_route_id ON trips (route_id);
CREATE INDEX idx_trips_service_id ON trips (service_id);
CREATE INDEX idx_trips_trip_id ON trips (trip_id);
