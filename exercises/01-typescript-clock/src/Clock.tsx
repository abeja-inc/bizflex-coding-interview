import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';
import timezone from 'dayjs/plugin/timezone';
import isToday from 'dayjs/plugin/isToday';
import isTomorrow from 'dayjs/plugin/isTomorrow';
import isYesterday from 'dayjs/plugin/isYesterday';
import './Clock.css';
import { useContext } from 'react';
import { ClockContext } from './ClockContext';

dayjs.extend(utc);
dayjs.extend(timezone);
dayjs.extend(isYesterday);
dayjs.extend(isToday);
dayjs.extend(isTomorrow);

export const Clock: React.VFC<{
  title: string;
  timeZone: string;
}> = ({ title, timeZone }) => {
  const { now, localTimeZone } = useContext(ClockContext);

  // Show difference between local time and the time in the specified time zone.
  const tzNow = now.tz(timeZone);
  const tzLocalNow = tzNow.tz(localTimeZone, true);
  const day = tzLocalNow.isYesterday()
    ? '昨日'
    : tzLocalNow.isToday()
    ? '今日'
    : tzLocalNow.isTomorrow()
    ? '明日'
    : '';
  const hours = tzLocalNow.diff(now, 'hour');

  return (
    <div className="Clock">
      <h2>
        <div className="diff">
          {day}、{hours >= 0 && '+'}
          {hours}時間
        </div>
        <div className="title">{title}</div>
      </h2>
      <p>{tzNow.format('H:mm')}</p>
    </div>
  );
};
