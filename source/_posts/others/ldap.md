---
title: Java 接入LDAP小知识
date: 2018-12-13 10:14:11
toc: true
# thumbnail: https://avatars3.githubusercontent.com/u/7270177?s=460&v=4
tags:
  - ldap
  - java
categories:
  - biz
---


```ini
ldap.url=ldap://10.1.3.129
```

```xml
<!--ldap configuration-->
<bean id="contextSource" class="org.springframework.ldap.core.support.LdapContextSource">
    <property name="url" value="${ldap.url}"/>
    <property name="base" value="ou=people,dc=qunhe,dc=cc"/>
    <property name="baseEnvironmentProperties">
        <map>
            <entry key="com.sun.jndi.ldap.connect.timeout" value="5000"/>
        </map>
    </property>
</bean>

<bean id="ldapTemplate" class="org.springframework.ldap.core.LdapTemplate">
    <constructor-arg ref="contextSource"/>
</bean>
```