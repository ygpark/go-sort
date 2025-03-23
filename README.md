# sorta

`sorta`는 리눅스의 `sort` 명령어를 Go 언어로 포팅한 경량 CLI 도구입니다. 텍스트 파일 또는 표준 입력을 받아 정렬된 결과를 출력하며, `-u` 옵션으로 중복된 줄을 제거할 수 있습니다.

## 특징

- 문자열 오름차순 정렬
- 중복 줄 제거 (`-u` 옵션)
- 파일 입력 또는 표준 입력 처리

## 설치

Go가 설치되어 있다면 아래 명령어로 빌드할 수 있습니다:

```bash
go build -o sorta
```

## 사용법

### 파일을 정렬

```bash
sorta filename.txt
```

### 중복 제거하며 정렬

```bash
sorta -u filename.txt
```

### 파이프를 이용한 정렬

```bash
cat filename.txt | sorta
```

### 파이프 + 중복 제거

```bash
cat filename.txt | sorta -u
```

## 예시

```txt
$ cat sample.txt
banana
apple
apple
carrot

$ sorta sample.txt
apple
apple
banana
carrot

$ sorta -u sample.txt
apple
banana
carrot
```

## 향후 추가 예정 기능

- 정렬 기준 지정 (숫자, 대소문자 무시 등)
- 역순 정렬
- 필드 기반 정렬
- 안정 정렬 옵션

## 라이센스

MIT License
