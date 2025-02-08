import type { ComponentType, JSX } from 'react';

import { IndexPage } from '@/pages/IndexPage/IndexPage';
import { InitDataPage } from '@/pages/InitDataPage.tsx';
import { LaunchParamsPage } from '@/pages/LaunchParamsPage.tsx';

import {ServicePage} from '@/pages/ServicePage/ServicePage'


interface Route {
  path: string;
  Component: ComponentType;
  title?: string;
  icon?: JSX.Element;
}

export const routes: Route[] = [
  { path: '/', Component: IndexPage },
  { path: '/init-data', Component: InitDataPage, title: 'Init Data' },
  { path: '/servicePage', Component: ServicePage, title: 'Записаться' },
  { path: '/launch-params', Component: LaunchParamsPage, title: 'Launch Params' },
];
