## OrientDB
OrientDB is a multi-model NoSQL database management system with native support for graphs, documents, full text, etc. and has a special focus on transaction performance.

### Setup
IF you want to expose OrientDB to `0.0.0.0`, you have to manually configure this in the OrientDB config file (default path: './volumes/orientdb-config/orientdb-server-config.xml').

Replace this:
```xml
<network>
    <protocols>
        <protocol implementation="com.orientechnologies.orient.server.network.protocol.binary.ONetworkProtocolBinary" name="binary"/>
    </protocols>
    <listeners>
        <listener protocol="binary" socket="default" port-range="2424-2430" ip-address="127.0.0.1"/>
    </listeners>
</network>
```

with this:
```xml
<network>
    <protocols>
        <protocol implementation="com.orientechnologies.orient.server.network.protocol.binary.ONetworkProtocolBinary" name="binary"/>
        <protocol implementation="com.orientechnologies.orient.server.network.protocol.http.ONetworkProtocolHttpDb" name="http"/>
    </protocols>
    <listeners>
        <listener protocol="binary" socket="default" port-range="2424-2430" ip-address="0.0.0.0"/>
        <listener protocol="http" ip-address="0.0.0.0" port-range="2480-2490">
            <parameters>
                <parameter name="network.http.charset" value="utf-8"/>
            </parameters>
            <commands>
                <command
                        pattern="GET|www GET|studio/ GET| GET|*.htm GET|*.html GET|*.xml GET|*.jpeg GET|*.jpg GET|*.png GET|*.gif GET|*.js GET|*.css GET|*.swf GET|*.ico GET|*.txt GET|*.otf GET|*.pjs GET|*.svg"
                        implementation="com.orientechnologies.orient.server.network.protocol.http.command.get.OServerCommandGetStaticContent">
                    <parameters>
                        <entry name="http.cache:*.htm *.html" value="Cache-Control: no-cache, no-store, max-age=0, must-revalidate\r\nPragma: no-cache"/>
                        <entry name="http.cache:default" value="Cache-Control: max-age=120"/>
                    </parameters>
                </command>
            </commands>
        </listener>
    </listeners>
</network>
```