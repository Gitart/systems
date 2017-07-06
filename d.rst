=======================================
Описание процедуры
=======================================

.. image:: https://travis-ci.org/jd/daiquiri.png?branch=master
    :target: https://travis-ci.org/jd/daiquiri
    :alt: Build Status

.. image:: https://img.shields.io/pypi/v/daiquiri.svg
    :target: https://pypi.python.org/pypi/daiquiri
    :alt: Latest Version

Описание новостных лент 

* Free software: Apache license
* Source: https://github.com/jd/daiquiri

Инсталяция
============

  pip install daiquiri

Если использование JSON
  pip install daiquiri[json]

If you want to eanble systemd support, you must install the `systemd` flavor::

  pip install daiquiri[systemd]

Пример
=====

Новость

.. literalinclude:: ../../examples/basic.py

Новость пример

.. literalinclude:: ../../examples/output.py


Picking format
--------------

Пример написать просомтр например

You can provide any class of type `logging.Formatter` as a formatter.

.. literalinclude:: ../../examples/formatter.py

Python warning support
----------------------

The Python `warnings` module is sometimes used by applications and library to
emit warnings. By default, they are printed on `stderr`. Daiquiri overrides
this by default and log warnings to the `py.warnings` logger.

This can be disabled by passing the `capture_warnings=False` argument to
`daiquiri.setup`.

Extra usage
-----------

While it's not mandatory to use `daiquiri.getLogger` to get a logger instead of
`logging.getLogger`, it is recommended as daiquiri provides an enhanced version
of the logger object. It allows any keyword argument to be passed to the
logging method and that will be available as part of the record.

.. literalinclude:: ../../examples/extra.py


Syslog support
--------------

The `daiquiri.output.Syslog` output provides syslog output.


Systemd journal support
-----------------------

The `daiquiri.output.Journal` output provides systemd journal support. All the
extra argument passed to the logger will be shipped as extra keys to the
journal.


File support
------------

The `daiquiri.output.File` output class provides support to log into a file.
