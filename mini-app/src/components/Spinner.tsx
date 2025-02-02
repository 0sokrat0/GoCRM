// components/Spinner.tsx
import { FC } from 'react';
import styles from './Spinner.module.css';

export const Spinner: FC = () => (
  <div className={styles.spinner}>
    <div className={styles.dot1}></div>
    <div className={styles.dot2}></div>
  </div>
);