import React from 'react';
import dayjs, { Dayjs } from 'dayjs';
import utc from 'dayjs/plugin/utc';
import timezone from 'dayjs/plugin/timezone';

dayjs.extend(utc);
dayjs.extend(timezone);

export type ClockState = {
  // The local time zone. The default is the user's time zone.
  localTimeZone: string;

  // The current date and time in the `localTimeZone`.
  now: Dayjs;

  // The total number of updates since the start up.
  updates: number;

  // Update interval in milliseconds. Stop timer if the value is `null`.
  refreshInterval: number | null;
};

export const DEFAULT_CLOCK_STATE: ClockState = {
  localTimeZone: 'Asia/Tokyo',
  now: dayjs.unix(0),
  updates: 0,
  refreshInterval: 1000,
};

export const ClockContext = React.createContext(DEFAULT_CLOCK_STATE);

export const ClockProvider: React.VFC<{
  localTimeZone?: string;
  // Update interval in milliseconds. `1000` by default.
  refreshInterval?: number | null | undefined;
  children: React.ReactNode;
}> = ({ children, refreshInterval = 1000, localTimeZone = 'Asia/Tokyo' }) => {
  const value: ClockState = {
    localTimeZone,
    now: dayjs().tz(localTimeZone),
    updates: 0,
    refreshInterval,
  };

  return <ClockContext.Provider value={value}>{children}</ClockContext.Provider>;
};
