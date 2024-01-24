# Log Agent

Read your logs in real time and send them to a remote server.

```yaml
input:
  - name: nginx-access-log
    path: access.log
    uses: tail
transformer:
    - for_input: nginx-access-log
      to_output: stdout
      uses: json-transformer
output:
  - name: stdout
    stdout: true
```