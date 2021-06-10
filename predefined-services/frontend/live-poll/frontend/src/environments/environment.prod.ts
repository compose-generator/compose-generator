/*
 * Copyright © Live-Poll 2020-2021. All rights reserved
 */

export const environment = {
    name: 'production',
    production: true,
    apiBaseWebsocketUrl: 'wss://${{LIVE_POLL_API_URL}}/v1',
    apiBaseUrl: 'https://${{LIVE_POLL_API_URL}}/v1',
    useSecureCookies: true,
    cookieConsentUrl: 'www.live-poll.de'
  };
  