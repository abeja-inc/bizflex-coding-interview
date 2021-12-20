import { useContext } from 'react';
import { ClockContext } from './ClockContext';
import './Settings.css';

export const Settings: React.VFC<{ onIntervalChange?: (interval: number | null) => void }> = ({
  onIntervalChange,
}) => {
  const { updates, refreshInterval, now } = useContext(ClockContext);

  return (
    <div className="Settings">
      Updated {updates} times / {now.format('HH:mm:ss')} / Interval:
      <select
        value={refreshInterval ?? ''}
        onChange={(e) => {
          const value = e.target.value === '' ? null : Number(e.target.value);
          const interval = Number.isNaN(value) ? null : value;

          onIntervalChange && onIntervalChange(interval);
        }}>
        <option value="">stop timer</option>
        <option value="100">100 msec</option>
        <option value="500">500 msec</option>
        <option value="1000">1 sec</option>
        <option value="5000">5 sec</option>
      </select>
    </div>
  );
};
