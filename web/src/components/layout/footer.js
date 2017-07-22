import React from 'react'
import { footerText } from 'config'
import styles from './footer.less'

const Footer = () => (
  <div className={styles.footer}>
    {footerText}
  </div>
)

export default Footer
